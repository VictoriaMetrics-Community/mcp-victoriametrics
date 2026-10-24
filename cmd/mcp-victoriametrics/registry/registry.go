package registry

import (
	"github.com/mark3labs/mcp-go/server"

	"github.com/VictoriaMetrics/mcp-victoriametrics/cmd/mcp-victoriametrics/config"
	"github.com/VictoriaMetrics/mcp-victoriametrics/cmd/mcp-victoriametrics/prompts"
	"github.com/VictoriaMetrics/mcp-victoriametrics/cmd/mcp-victoriametrics/tools"
)

// Register registers tools and prompts.
func Register(s *server.MCPServer, c *config.Config) {
	tools.Register(s, c)
	prompts.Register(s, c)
}
