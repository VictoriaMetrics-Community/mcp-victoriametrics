package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics/config"
)

var (
	toolMetrics = mcp.NewTool("metrics",
		mcp.WithDescription("List of available metrics of the VictoriaMetrics instance. This tool uses `/api/v1/label/__name__/values` endpoint of VictoriaMetrics API."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "List of metric names",
			ReadOnlyHint:    true,
			DestructiveHint: false,
			OpenWorldHint:   true,
		}),
		mcp.WithString("tenant",
			mcp.Title("Tenant name"),
			mcp.Description("Name of the tenant for which the list of metrics will be displayed"),
			mcp.DefaultString("0"),
			mcp.Pattern(`^([0-9]+)(\:[0-9]+)?$`),
		),
		mcp.WithString("match",
			mcp.Title("Match series for metric names"),
			mcp.Description("Time series selector argument that selects the series from which to read the metrics"),
			mcp.DefaultString(""),
		),
		mcp.WithString("start",
			mcp.Title("Start timestamp"),
			mcp.Description("Start timestamp for selection metric names"),
			mcp.DefaultString(""),
			mcp.Pattern(`^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)|([0-9]+)$`),
		),
		mcp.WithString("end",
			mcp.Title("End timestamp"),
			mcp.Description("End timestamp for selection metric names"),
			mcp.DefaultString(""),
			mcp.Pattern(`^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)|([0-9]+)$`),
		),
		mcp.WithNumber("limit",
			mcp.Title("Maximum number of metric names"),
			mcp.Description("Maximum number of metric names to return"),
			mcp.DefaultNumber(0),
			mcp.Min(0),
		),
	)
)

func toolMetricsHandler(ctx context.Context, cfg *config.Config, tcr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	tenant, err := GetToolReqParam[string](tcr, "tenant", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	match, err := GetToolReqParam[string](tcr, "match", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	start, err := GetToolReqParam[string](tcr, "start", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	end, err := GetToolReqParam[string](tcr, "end", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	limit, err := GetToolReqParam[float64](tcr, "limit", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return getLabelValues(ctx, cfg, tenant, "__name__", match, start, end, limit)
}

func RegisterToolMetrics(s *server.MCPServer, c *config.Config) {
	s.AddTool(toolMetrics, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolMetricsHandler(ctx, c, request)
	})
}
