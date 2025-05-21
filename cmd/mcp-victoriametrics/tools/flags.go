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
	toolFlags = mcp.NewTool("flags",
		mcp.WithDescription("List of non-default flags (parameters) of the VictoriaMetrics instance. This tools uses `/flags` endpoint of VictoriaMetrics API."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "List of non-default flags (parameters)",
			ReadOnlyHint:    ptr(true),
			DestructiveHint: ptr(false),
			OpenWorldHint:   ptr(true),
		}),
	)
)

func toolFlagsHandler(ctx context.Context, cfg *config.Config, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.AdminAPIURL("flags"), nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create request: %v", err)), nil
	}
	return GetTextBodyForRequest(req, cfg), nil
}

func RegisterToolFlags(s *server.MCPServer, c *config.Config) {
	s.AddTool(toolFlags, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolFlagsHandler(ctx, c, request)
	})
}
