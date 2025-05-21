package tools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics/config"
)

var (
	toolTenants = mcp.NewTool("tenants",
		mcp.WithDescription("List of tenants of the VictoriaMetrics instance.  This tool uses `/admin/tenants` endpoint of VictoriaMetrics API."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "List of tenants",
			ReadOnlyHint:    ptr(true),
			DestructiveHint: ptr(false),
			OpenWorldHint:   ptr(true),
		}),
	)
)

func toolTenantsHandler(ctx context.Context, cfg *config.Config, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.AdminAPIURL("admin", "tenants"), nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create request: %v", err)), nil
	}
	return GetTextBodyForRequest(req, cfg), nil
}

func RegisterToolTenants(s *server.MCPServer, c *config.Config) {
	s.AddTool(toolTenants, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolTenantsHandler(ctx, c, request)
	})
}
