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
	toolLabelsValues = mcp.NewTool("label_values",
		mcp.WithDescription("List of label values from VictoriaMetrics instance for a provided label name. This tool uses `/api/v1/label/{labelName}/values` endpoint of VictoriaMetrics API."),
		mcp.WithString("tenant",
			mcp.Title("Tenant name"),
			mcp.Description("Name of the tenant for which the list of label values will be displayed"),
			mcp.DefaultString("0"),
			mcp.Pattern(`^([0-9]+)(\:[0-9]+)?$`),
		),
		mcp.WithString("label_name",
			mcp.Required(),
			mcp.Title("Label name"),
			mcp.Description("Name of the label to query its values"),
			mcp.Pattern(`^.+$`),
		),
		mcp.WithString("match",
			mcp.Title("Match series for labels values"),
			mcp.Description("Time series selector argument that selects the series from which to read the label values"),
			mcp.DefaultString(""),
		),
		mcp.WithString("start",
			mcp.Title("Start timestamp"),
			mcp.Description("Start timestamp for selection labels values"),
			mcp.DefaultString(""),
			mcp.Pattern(`^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)|([0-9]+)$`),
		),
		mcp.WithString("end",
			mcp.Title("End timestamp"),
			mcp.Description("End timestamp for selection labels values"),
			mcp.DefaultString(""),
			mcp.Pattern(`^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)|([0-9]+)$`),
		),
		mcp.WithNumber("limit",
			mcp.Title("Maximum number of label values"),
			mcp.Description("Maximum number of label values to return"),
			mcp.DefaultNumber(0),
			mcp.Min(0),
		),
	)
)

func toolLabelValuesHandler(ctx context.Context, cfg *config.Config, tcr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	tenant, err := GetToolReqParam[string](tcr, "tenant", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	match, err := GetToolReqParam[string](tcr, "match", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	labelName, err := GetToolReqParam[string](tcr, "label_name", true)
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

	return getLabelValues(ctx, cfg, tenant, labelName, match, start, end, limit)
}

func getLabelValues(ctx context.Context, cfg *config.Config, tenant string, labelName string, match string, start string, end string, limit float64) (*mcp.CallToolResult, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cfg.SelectAPIURL(tenant, "api", "v1", "label", labelName, "values"), nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create request: %v", err)), nil
	}

	q := req.URL.Query()
	if match != "" {
		q.Add("match[]", match)
	}
	if start != "" {
		q.Add("start", start)
	}
	if end != "" {
		q.Add("end", end)
	}
	if limit != 0 {
		q.Add("limit", fmt.Sprintf("%.f", limit))
	}
	req.URL.RawQuery = q.Encode()

	return GetTextBodyForRequest(req, cfg), nil
}

func RegisterToolLabelValues(s *server.MCPServer, c *config.Config) {
	s.AddTool(toolLabelsValues, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolLabelValuesHandler(ctx, c, request)
	})
}
