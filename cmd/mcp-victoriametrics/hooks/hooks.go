package hooks

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/VictoriaMetrics/metrics"
)

func New(ms *metrics.Set) *server.Hooks {
	hooks := &server.Hooks{}

	hooks.AddAfterInitialize(func(_ context.Context, _ any, message *mcp.InitializeRequest, _ *mcp.InitializeResult) {
		ms.GetOrCreateCounter(fmt.Sprintf(
			`mcp_victoriametrics_initialize_total{client_name="%s",client_version="%s"}`,
			message.Params.ClientInfo.Name,
			message.Params.ClientInfo.Version,
		)).Inc()
	})

	hooks.AddAfterListTools(func(_ context.Context, _ any, _ *mcp.ListToolsRequest, _ *mcp.ListToolsResult) {
		ms.GetOrCreateCounter(`mcp_victoriametrics_list_tools_total`).Inc()
	})

	hooks.AddAfterListResources(func(_ context.Context, _ any, _ *mcp.ListResourcesRequest, _ *mcp.ListResourcesResult) {
		ms.GetOrCreateCounter(`mcp_victoriametrics_list_resources_total`).Inc()
	})

	hooks.AddAfterListPrompts(func(_ context.Context, _ any, _ *mcp.ListPromptsRequest, _ *mcp.ListPromptsResult) {
		ms.GetOrCreateCounter(`mcp_victoriametrics_list_prompts_total`).Inc()
	})

	hooks.AddAfterCallTool(func(_ context.Context, _ any, message *mcp.CallToolRequest, result *mcp.CallToolResult) {
		ms.GetOrCreateCounter(fmt.Sprintf(
			`mcp_victoriametrics_call_tool_total{name="%s",is_error="%t"}`,
			message.Params.Name,
			result.IsError,
		)).Inc()
	})

	hooks.AddAfterGetPrompt(func(_ context.Context, _ any, message *mcp.GetPromptRequest, _ *mcp.GetPromptResult) {
		ms.GetOrCreateCounter(fmt.Sprintf(
			`mcp_victoriametrics_get_prompt_total{name="%s"}`,
			message.Params.Name,
		)).Inc()
	})

	hooks.AddAfterReadResource(func(_ context.Context, _ any, message *mcp.ReadResourceRequest, _ *mcp.ReadResourceResult) {
		ms.GetOrCreateCounter(fmt.Sprintf(
			`mcp_victoriametrics_read_resource_total{name="%s"}`,
			message.Params.URI,
		)).Inc()
	})

	hooks.AddOnError(func(_ context.Context, _ any, method mcp.MCPMethod, _ any, err error) {
		ms.GetOrCreateCounter(fmt.Sprintf(
			`mcp_victoriametrics_error_total{method="%s",error="%s"}`,
			method,
			err,
		)).Inc()
	})

	return hooks
}
