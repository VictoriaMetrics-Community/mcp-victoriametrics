package tools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/VictoriaMetrics-Community/mcp-victoriametrics/cmd/mcp-victoriametrics/config"
)

const toolNameRules = "rules"

func toolRules(c *config.Config) mcp.Tool {
	options := []mcp.ToolOption{
		mcp.WithDescription("List of alerting and recording rules of VictoriaMetrics instance. This tool uses `/api/v1/rules` endpoint of vmalert API proxied by VictoriaMetrics API."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:           "List of alerting and recording rules",
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
		options = append(
			options,
			mcp.WithString("tenant",
				mcp.Title("Tenant name"),
				mcp.Description("Name of the tenant for which the list of rules will be displayed"),
				mcp.DefaultString("0"),
				mcp.Pattern(`^([0-9]+)(\:[0-9]+)?$`),
			),
		)
	}
	options = append(
		options,
		mcp.WithString("type",
			mcp.Title("Rules type"),
			mcp.Description("Rules type to be displayed: alert or record"),
			mcp.DefaultString(""),
			mcp.Enum("alert", "record"),
		),
		mcp.WithString("filter",
			mcp.Title("Extra filter for rules"),
			mcp.Description("Extra filter for rules with possible problems: unhealthy (rules that get some errors during evaluation) or noMatch (rules that don't match any time series)"),
			mcp.DefaultString(""),
			mcp.Enum("unhealthy", "noMatch"),
		),
		mcp.WithBoolean("exclude_alerts",
			mcp.Title("Exclude alerts"),
			mcp.Description("Exclude alerts from the list"),
			mcp.DefaultBool(false),
		),
		mcp.WithArray("rule_names",
			mcp.Title("Rule names"),
			mcp.Description("Filter rules by name"),
			mcp.Items(map[string]any{"type": "string"}),
		),
		mcp.WithArray("rule_groups",
			mcp.Title("Rule groups"),
			mcp.Description("Filter rules by group names"),
			mcp.Items(map[string]any{"type": "string"}),
		),
		mcp.WithArray("rule_files",
			mcp.Title("Rule files"),
			mcp.Description("Filter rules by file names"),
			mcp.Items(map[string]any{"type": "string"}),
		),
	)
	return mcp.NewTool(toolNameRules, options...)
}

func toolRulesHandler(ctx context.Context, cfg *config.Config, tcr mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ruleType, err := GetToolReqParam[string](tcr, "type", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	filter, err := GetToolReqParam[string](tcr, "filter", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	excludeAlerts, err := GetToolReqParam[bool](tcr, "exclude_alerts", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	ruleNames, err := GetToolReqParam[[]string](tcr, "rule_names", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	ruleGroups, err := GetToolReqParam[[]string](tcr, "rule_groups", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	ruleFiles, err := GetToolReqParam[[]string](tcr, "rule_files", false)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	req, err := CreateSelectRequest(ctx, cfg, tcr, "vmalert", "api", "v1", "rules")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create request: %v", err)), nil
	}

	q := req.URL.Query()
	if ruleType != "" {
		q.Add("type", ruleType)
	}
	if filter != "" {
		q.Add("filter", filter)
	}
	if excludeAlerts {
		q.Add("exclude_alerts", "true")
	}
	for _, ruleName := range ruleNames {
		q.Add("rule_names", ruleName)
	}
	for _, ruleGroup := range ruleGroups {
		q.Add("rule_groups", ruleGroup)
	}
	for _, ruleFile := range ruleFiles {
		q.Add("rule_files", ruleFile)
	}
	req.URL.RawQuery = q.Encode()

	return GetTextBodyForRequest(req, cfg), nil
}

func RegisterToolRules(s *server.MCPServer, c *config.Config) {
	if c.IsToolDisabled(toolNameRules) {
		return
	}
	s.AddTool(toolRules(c), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return toolRulesHandler(ctx, c, request)
	})
}
