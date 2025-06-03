package tools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics/config"
)

const toolNameTopQueries = "top_queries"

func toolTopQueries(c *config.Config) mcp.Tool {
	options := []mcp.ToolOption{
		mcp.WithDescription(`Top queries.
This tool can determine top queries of the following query types:
- the most frequently executed queries;
- queries with the biggest average execution duration;
- queries that took the most summary time for execution.
This information is obtained from the "/api/v1/status/top_queries" HTTP endpoint of VictoriaMetrics API.
`),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "Top queries",
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
				mcp.Description("Name of the tenant for which the top queries will be displayed"),
				mcp.DefaultString("0"),
				mcp.Pattern(`^([0-9]+)(\:[0-9]+)?$`),
			),
		)
	}
	options = append(
		options,
		mcp.WithNumber("topN",
			mcp.Required(),
			mcp.Title("Top N"),
			mcp.Description("The number of top entries to return in the response. By default is 20."),
			mcp.DefaultNumber(20),
			mcp.Min(1),
		),
		mcp.WithString("maxLifetime",
			mcp.Title("Max lifetime"),
			mcp.Description("Max lifetime of the queries to be taken into account during stats calculation. By default is 10m."),
			mcp.DefaultString("10m"),
			mcp.Pattern(`^([0-9]+)([a-z]+)$`),
		),
	)
	return mcp.NewTool(toolNameTopQueries, options...)
}

func toolTopQueriesHandler(ctx context.Context, cfg *config.Config, tcr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	tenant, err := GetToolReqParam[string](tcr, "tenant", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	topN, err := GetToolReqParam[float64](tcr, "topN", true)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	if topN < 1 {
		topN = 20
	}

	maxLifetime, err := GetToolReqParam[string](tcr, "maxLifetime", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.SelectAPIURL(tenant, "api", "v1", "status", "top_queries"), nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create request: %v", err)), nil
	}

	query := req.URL.Query()
	query.Set("topN", fmt.Sprintf("%d", int(topN)))
	if maxLifetime != "" {
		query.Set("focusLabel", maxLifetime)
	}
	req.URL.RawQuery = query.Encode()

	return GetTextBodyForRequest(req, cfg), nil
}

func RegisterToolTopQueries(s *server.MCPServer, c *config.Config) {
	if c.IsToolDisabled(toolNameTopQueries) {
		return
	}
	s.AddTool(toolTopQueries(c), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolTopQueriesHandler(ctx, c, request)
	})
}
