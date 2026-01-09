package logging

import (
	"fmt"

	"github.com/mark3labs/mcp-go/util"
)

// Ensure MCPLoggerAdapter implements util.Logger
var _ util.Logger = (*MCPLoggerAdapter)(nil)

// MCPLoggerAdapter adapts our Logger to mcp-go's util.Logger interface
type MCPLoggerAdapter struct {
	logger *Logger
}

// NewMCPLoggerAdapter creates a new adapter for mcp-go util.Logger interface
func NewMCPLoggerAdapter(l *Logger) *MCPLoggerAdapter {
	return &MCPLoggerAdapter{logger: l}
}

// Infof implements util.Logger.Infof
func (a *MCPLoggerAdapter) Infof(format string, v ...any) {
	if a.logger.IsEnabled() {
		a.logger.Info(fmt.Sprintf(format, v...))
	}
}

// Errorf implements util.Logger.Errorf
func (a *MCPLoggerAdapter) Errorf(format string, v ...any) {
	if a.logger.IsEnabled() {
		a.logger.Error(fmt.Sprintf(format, v...))
	}
}
