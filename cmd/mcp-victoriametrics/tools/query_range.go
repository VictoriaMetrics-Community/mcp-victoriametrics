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
	toolQueryRange = mcp.NewTool("query_range",
		mcp.WithDescription("Range query executes the query expression at the given [start…end] time range with the given step. The result of Range query is a list of time series matching the filter in query expression. Each returned series contains (timestamp, value) results for the query executed at start, start+step, start+2*step, …, start+N*step timestamps. In other words, Range query is an Instant query executed independently at start, start+step, …, start+N*step timestamps with the only difference that an instant query does not return ephemeral samples (see below). Instead, if the database does not contain any samples for the requested time and step, it simply returns an empty result. This tool uses `/api/v1/query_range` endpoint of VictoriaMetrics API."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "Range Query",
			ReadOnlyHint:    true,
			DestructiveHint: false,
			OpenWorldHint:   true,
		}),
		mcp.WithString("tenant",
			mcp.Title("Tenant name"),
			mcp.Description("Name of the tenant for which the data will be displayed"),
			mcp.DefaultString("0"),
			mcp.Pattern(`^([0-9]+)(\:[0-9]+)?$`),
		),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Title("MetricsQL or PromQL expression"),
			mcp.Description("MetricsQL or PromQL expression for the query of the data"),
		),
		mcp.WithString("start",
			mcp.Required(),
			mcp.Title("Start timestamp"),
			mcp.Description("The starting timestamp of the time range for query evaluation"),
			mcp.Pattern(`^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)|([0-9]+)$`),
		),
		mcp.WithString("end",
			mcp.Title("End timestamp"),
			mcp.Description("The ending timestamp of the time range for query evaluation. If the end isn’t set, then the end is automatically set to the current time."),
			mcp.Pattern(`^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)|([0-9]+)$`),
		),
		mcp.WithString("step",
			mcp.Title("Step"),
			mcp.Description("the interval between data points, which must be returned from the range query. The query is executed at start, start+step, start+2*step, …, start+N*step timestamps, where N is the whole number of steps that fit between start and end. end is included only when it equals to start+N*step. If the step isn’t set, then it default to 5m (5 minutes)."),
			mcp.Pattern(`^([0-9]+)([a-z]+)$`),
		),
		mcp.WithString("timeout",
			mcp.Title("Timeout"),
			mcp.Description("optional query timeout. For example, timeout=5s. Query is canceled when the timeout is reached. By default the timeout is set to the value of -search.maxQueryDuration command-line flag passed to single-node VictoriaMetrics or to vmselect component in VictoriaMetrics cluster."),
			mcp.Pattern(`^([0-9]+)([a-z]+)$`),
		),
	)
)

func toolQuerysRangeHandler(ctx context.Context, cfg *config.Config, tcr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	tenant, err := GetToolReqParam[string](tcr, "tenant", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	query, err := GetToolReqParam[string](tcr, "query", true)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	start, err := GetToolReqParam[string](tcr, "start", true)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	end, err := GetToolReqParam[string](tcr, "end", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	step, err := GetToolReqParam[string](tcr, "step", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	timeout, err := GetToolReqParam[string](tcr, "timeout", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.SelectAPIURL(tenant, "api", "v1", "query"), nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create request: %v", err)), nil
	}

	q := req.URL.Query()
	q.Add("query", query)
	q.Add("start", start)
	if end != "" {
		q.Add("end", end)
	}
	if step != "" {
		q.Add("step", step)
	}
	if timeout != "" {
		q.Add("timeout", step)
	}
	req.URL.RawQuery = q.Encode()

	return GetTextBodyForRequest(req, cfg), nil
}

func RegisterToolQueryRange(s *server.MCPServer, c *config.Config) {
	s.AddTool(toolQueryRange, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolQuerysRangeHandler(ctx, c, request)
	})
}
