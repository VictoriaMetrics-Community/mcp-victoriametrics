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
	toolExport = mcp.NewTool("export",
		mcp.WithDescription("Export time series to file (json or csv) from VictoriaMetrics instance. This tool uses `/api/v1/export` or `/api/v1/export/csv` endpoints of VictoriaMetrics API)"),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "Export time series",
			ReadOnlyHint:    true,
			DestructiveHint: false,
			OpenWorldHint:   true,
		}),
		mcp.WithString("tenant",
			mcp.Title("Tenant name"),
			mcp.Description("Name of the tenant for which the data will be exported"),
			mcp.DefaultString("0"),
			mcp.Pattern(`^([0-9]+)(\:[0-9]+)?$`),
		),
		mcp.WithString("match",
			mcp.Required(),
			mcp.Title("Match series for export"),
			mcp.Description("Time series selector argument that selects the series for export"),
		),
		mcp.WithString("start",
			mcp.Title("Start timestamp"),
			mcp.Description("Start timestamp for export"),
			mcp.DefaultString(""),
			mcp.Pattern(`^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)|([0-9]+)$`),
		),
		mcp.WithString("end",
			mcp.Title("End timestamp"),
			mcp.Description("End timestamp for export"),
			mcp.DefaultString(""),
			mcp.Pattern(`^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)|([0-9]+)$`),
		),
		mcp.WithString("format",
			mcp.Required(),
			mcp.Description("Export format: json (default) or csv"),
			mcp.DefaultString("json"),
			mcp.Enum("csv", "json"),
		),
	)
)

func toolExportHandler(ctx context.Context, cfg *config.Config, tcr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	tenant, err := GetToolReqParam[string](tcr, "tenant", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	match, err := GetToolReqParam[string](tcr, "match", true)
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

	format, err := GetToolReqParam[string](tcr, "format", true)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	var req *http.Request

	switch format {
	case "json":
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, cfg.SelectAPIURL(tenant, "api", "v1", "export"), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("failed to create request: %v", err)), nil
		}
	case "csv":
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, cfg.SelectAPIURL(tenant, "api", "v1", "export", "csv"), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("failed to create request: %v", err)), nil
		}
	default:
		return mcp.NewToolResultError(fmt.Sprintf("unsupported format: %s", format)), nil
	}

	q := req.URL.Query()
	q.Add("match[]", match)
	if start != "" {
		q.Add("start", start)
	}
	if end != "" {
		q.Add("end", end)
	}
	q.Add("format", format)
	req.URL.RawQuery = q.Encode()

	return GetTextBodyForRequest(req, cfg), nil
}

func RegisterToolExport(s *server.MCPServer, c *config.Config) {
	s.AddTool(toolExport, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolExportHandler(ctx, c, request)
	})
}
