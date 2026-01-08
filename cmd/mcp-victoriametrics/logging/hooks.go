package logging

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// NewHooks creates MCP server hooks for logging
func (l *Logger) NewHooks() *server.Hooks {
	if !l.enabled {
		return &server.Hooks{}
	}

	hooks := &server.Hooks{}

	hooks.AddOnRegisterSession(func(_ context.Context, session server.ClientSession) {
		l.Info("Session registered",
			"session_id", session.SessionID(),
		)
	})

	hooks.AddOnUnregisterSession(func(_ context.Context, session server.ClientSession) {
		l.Info("Session unregistered",
			"session_id", session.SessionID(),
		)
	})

	hooks.AddBeforeAny(func(ctx context.Context, id any, method mcp.MCPMethod, message any) {
		sessionID := extractSessionID(ctx)
		l.Info("MCP request received",
			"request_id", id,
			"session_id", sessionID,
			"method", string(method),
			"message", message,
		)
	})

	hooks.AddOnSuccess(func(ctx context.Context, id any, method mcp.MCPMethod, message any, result any) {
		sessionID := extractSessionID(ctx)
		l.Info("MCP request succeeded",
			"request_id", id,
			"session_id", sessionID,
			"method", string(method),
			"message", message,
			"result", result,
		)
	})

	hooks.AddOnError(func(ctx context.Context, id any, method mcp.MCPMethod, message any, err error) {
		sessionID := extractSessionID(ctx)
		l.Error("MCP request failed",
			"request_id", id,
			"session_id", sessionID,
			"method", string(method),
			"message", message,
			"error", err.Error(),
		)
	})

	hooks.AddAfterInitialize(func(_ context.Context, id any, msg *mcp.InitializeRequest, _ *mcp.InitializeResult) {
		l.Info("Client initialized",
			"request_id", id,
			"client_name", msg.Params.ClientInfo.Name,
			"client_version", msg.Params.ClientInfo.Version,
			"protocol_version", msg.Params.ProtocolVersion,
		)
	})

	hooks.AddAfterCallTool(func(_ context.Context, id any, msg *mcp.CallToolRequest, result *mcp.CallToolResult) {
		l.Info("Tool called",
			"request_id", id,
			"tool_name", msg.Params.Name,
			"is_error", result.IsError,
		)
	})

	return hooks
}

// extractSessionID extracts session ID from context
func extractSessionID(ctx context.Context) string {
	session := server.ClientSessionFromContext(ctx)
	if session != nil {
		return session.SessionID()
	}
	return ""
}
