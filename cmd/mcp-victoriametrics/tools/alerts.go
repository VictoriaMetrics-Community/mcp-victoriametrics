package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics/config"
)

const toolNameAlerts = "alerts"

func toolAlerts(c *config.Config) mcp.Tool {
	options := []mcp.ToolOption{
		mcp.WithDescription("List of firing and pending alerts of the VictoriaMetrics instance. This tool uses `/api/v1/alerts` endpoint of vmalert API, proxied by VictoriaMetrics API."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "List of alerts",
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
	if c.IsCluster() || c.IsCloud() {
		mcp.WithString("tenant",
			mcp.Title("Tenant name"),
			mcp.Description("Name of the tenant for which the list of alerts will be displayed"),
			mcp.DefaultString("0"),
			mcp.Pattern(`^([0-9]+)(\:[0-9]+)?$`),
		)
	}
	return mcp.NewTool(toolNameAlerts, options...)
}

func toolAlertsHandler(ctx context.Context, cfg *config.Config, tcr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	req, err := CreateSelectRequest(ctx, cfg, tcr, "vmalert", "api", "v1", "alerts")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create request: %v", err)), nil
	}
	return GetTextBodyForRequest(req, cfg), nil
}

func RegisterToolAlerts(s *server.MCPServer, c *config.Config) {
	if c.IsToolDisabled(toolNameAlerts) {
		return
	}
	s.AddTool(toolAlerts(c), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolAlertsHandler(ctx, c, request)
	})
}
