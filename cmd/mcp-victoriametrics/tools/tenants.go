package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/VictoriaMetrics/mcp-victoriametrics/cmd/mcp-victoriametrics/config"
)

const toolNameTenants = "tenants"

func toolTenants(c *config.Config) mcp.Tool {
	description := "List of tenants of the VictoriaMetrics instance.  This tool uses `/admin/tenants` endpoint of VictoriaMetrics API."
	if len(c.Tenants()) > 0 {
		description += " The result also contains aliases and descriptions of the tenants from the MCP server configuration (e.g. which environment each tenant belongs to)."
	}
	options := []mcp.ToolOption{
		mcp.WithDescription(description),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "List of tenants",
			ReadOnlyHint:    ptr(true),
			DestructiveHint: ptr(false),
			OpenWorldHint:   ptr(true),
		}),
	}
	if c.IsCloud() {
		options = append(
			options,
			mcp.WithString("deployment_id",
				mcp.Required(),
				mcp.Title("Deployment ID"),
				mcp.Description("Unique identifier of the deployment in VictoriaMetrics Cloud"),
				mcp.Pattern(`^[a-zA-Z0-9\-_]+$`),
			),
		)
	}
	return mcp.NewTool(toolNameTenants, options...)
}

func toolTenantsHandler(ctx context.Context, cfg *config.Config, tcr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	req, err := CreateAdminRequest(ctx, cfg, tcr, "admin", "tenants")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create request: %v", err)), nil
	}
	result := GetTextBodyForRequest(req, cfg)
	if len(cfg.Tenants()) > 0 {
		// Described tenants are returned even if the request has failed (e.g. /admin/tenants isn't allowed by vmauth)
		result.Content = append(result.Content, mcp.NewTextContent(
			"Tenants described in the MCP server configuration:\n"+describeTenants(cfg),
		))
	}
	return result, nil
}

// TenantsInstructions returns the instructions for AI assistant about the tenants described in VM_TENANTS_FILE
func TenantsInstructions(c *config.Config) string {
	if len(c.Tenants()) == 0 {
		return ""
	}
	return `
VictoriaMetrics cluster contains several tenants (e.g. for different environments), pass the ID of the right tenant to the "tenant" parameter of the tools.
The following tenants are described in the MCP server configuration:
` + describeTenants(c) + `If it's not clear which tenant (environment) the user is asking about, ask the user.
`
}

// describeTenants returns the list of tenants described in VM_TENANTS_FILE in a human-readable form
func describeTenants(c *config.Config) string {
	defaultTenant := c.DefaultTenantID()
	if id, err := config.CanonicalTenantID(defaultTenant); err == nil {
		defaultTenant = id
	}
	defaultTenantName := defaultTenant

	var sb strings.Builder
	for _, t := range c.Tenants() {
		sb.WriteString("- " + t.ID)
		if t.Alias != "" {
			sb.WriteString(" (alias: " + t.Alias + ")")
			if t.ID == defaultTenant {
				defaultTenantName = fmt.Sprintf("%s (alias: %s)", t.ID, t.Alias)
			}
		}
		if t.Description != "" {
			sb.WriteString(": " + strings.ReplaceAll(t.Description, "\n", "\n  "))
		}
		sb.WriteString("\n")
	}
	fmt.Fprintf(&sb, "If the \"tenant\" parameter is omitted, tenant %s is used.\n", defaultTenantName)
	return sb.String()
}

func RegisterToolTenants(s *server.MCPServer, c *config.Config) {
	if c.IsToolDisabled(toolNameTenants) {
		return
	}
	if !c.IsCluster() && !c.IsCloud() {
		return
	}
	s.AddTool(toolTenants(c), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolTenantsHandler(ctx, c, request)
	})
}
