package hooks

import (
	"context"
	"fmt"

	"github.com/VictoriaMetrics/metrics"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func New(ms *metrics.Set) *server.Hooks {
	hooks := &server.Hooks{}

	hooks.AddAfterInitialize(func(ctx context.Context, id any, message *mcp.InitializeRequest, _ *mcp.InitializeResult) {
		ms.GetOrCreateCounter(fmt.Sprintf(
			`mcp_victoriametrics_initialize_total{client_name="%s",client_version="%s"}`,
			message.Params.ClientInfo.Name,
			message.Params.ClientInfo.Version,
		)).Inc()
	})

	hooks.AddAfterListTools(func(ctx context.Context, id any, message *mcp.ListToolsRequest, result *mcp.ListToolsResult) {
		ms.GetOrCreateCounter(`mcp_victoriametrics_list_tools_total`).Inc()
	})

	hooks.AddAfterListResources(func(ctx context.Context, id any, message *mcp.ListResourcesRequest, result *mcp.ListResourcesResult) {
		ms.GetOrCreateCounter(`mcp_victoriametrics_list_resources_total`).Inc()
	})

	hooks.AddAfterListPrompts(func(ctx context.Context, id any, message *mcp.ListPromptsRequest, result *mcp.ListPromptsResult) {
		ms.GetOrCreateCounter(`mcp_victoriametrics_list_prompts_total`).Inc()
	})

	hooks.AddAfterCallTool(func(ctx context.Context, id any, message *mcp.CallToolRequest, result *mcp.CallToolResult) {
		ms.GetOrCreateCounter(fmt.Sprintf(
			`mcp_victoriametrics_call_tool_total{name="%s",is_error="%t"}`,
			message.Params.Name,
			result.IsError,
		)).Inc()
	})

	hooks.AddAfterGetPrompt(func(ctx context.Context, id any, message *mcp.GetPromptRequest, result *mcp.GetPromptResult) {
		ms.GetOrCreateCounter(fmt.Sprintf(
			`mcp_victoriametrics_get_prompt_total{name="%s"}`,
			message.Params.Name,
		)).Inc()
	})

	hooks.AddAfterReadResource(func(ctx context.Context, id any, message *mcp.ReadResourceRequest, result *mcp.ReadResourceResult) {
		ms.GetOrCreateCounter(fmt.Sprintf(
			`mcp_victoriametrics_read_resource_total{name="%s"}`,
			message.Params.URI,
		)).Inc()
	})

	hooks.AddOnError(func(ctx context.Context, id any, method mcp.MCPMethod, message any, err error) {
		ms.GetOrCreateCounter(fmt.Sprintf(
			`mcp_victoriametrics_error_total{method="%s",error="%s"}`,
			method,
			err,
		)).Inc()
	})

	return hooks
}
