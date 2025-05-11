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
	toolActiveQueries = mcp.NewTool("active_queries",
		mcp.WithDescription(`Active queries. This tool can determine currently active queries in the VictoriaMetrics instance.
This information is obtained from the "/api/v1/status/active_queries" HTTP endpoint of VictoriaMetrics API.`),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "Active queries",
			ReadOnlyHint:    true,
			DestructiveHint: false,
			OpenWorldHint:   true,
		}),
		mcp.WithString("tenant",
			mcp.Title("Tenant name"),
			mcp.Description("Name of the tenant for which the active queries will be displayed"),
			mcp.DefaultString("0"),
			mcp.Pattern(`^([0-9]+)(\:[0-9]+)?$`),
		),
	)
)

func toolActiveQueriesHandler(ctx context.Context, cfg *config.Config, tcr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	tenant, err := GetToolReqParam[string](tcr, "tenant", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.SelectAPIURL(tenant, "api", "v1", "status", "active_queries"), nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create request: %v", err)), nil
	}

	return GetTextBodyForRequest(req, cfg), nil
}

func RegisterToolActiveQueries(s *server.MCPServer, c *config.Config) {
	s.AddTool(toolActiveQueries, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolActiveQueriesHandler(ctx, c, request)
	})
}
