package tools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics/config"
)

const toolNamePrettifyQuery = "prettify_query"

func toolPrettifyQuery(c *config.Config) mcp.Tool {
	options := []mcp.ToolOption{
		mcp.WithDescription("Prettify (format) MetricsQL query. This tool uses `/prettify-query` endpoint of VictoriaMetrics API."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "Prettify Query",
			ReadOnlyHint:    ptr(true),
			DestructiveHint: ptr(false),
			OpenWorldHint:   ptr(true),
		}),
	}
	if c.IsCluster() {
		options = append(
			options,
			mcp.WithString("tenant",
				mcp.Title("Tenant name"),
				mcp.Description("Name of the tenant for which the data will be displayed"),
				mcp.DefaultString("0"),
				mcp.Pattern(`^([0-9]+)(\:[0-9]+)?$`),
			),
		)
	}
	options = append(
		options,
		mcp.WithString("query",
			mcp.Required(),
			mcp.Title("MetricsQL or PromQL expression"),
			mcp.Description(`MetricsQL or PromQL expression for prettification. This is the query that will be formatted.`),
		),
	)
	return mcp.NewTool(toolNamePrettifyQuery, options...)
}

func toolPrettifyQueryHandler(ctx context.Context, cfg *config.Config, tcr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	tenant, err := GetToolReqParam[string](tcr, "tenant", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	query, err := GetToolReqParam[string](tcr, "query", true)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.SelectAPIURL(tenant, "prettify-query"), nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create request: %v", err)), nil
	}

	q := req.URL.Query()
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()

	return GetTextBodyForRequest(req, cfg), nil
}

func RegisterToolPrettifyQuery(s *server.MCPServer, c *config.Config) {
	if c.IsToolDisabled(toolNamePrettifyQuery) {
		return
	}
	s.AddTool(toolPrettifyQuery(c), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolPrettifyQueryHandler(ctx, c, request)
	})
}
