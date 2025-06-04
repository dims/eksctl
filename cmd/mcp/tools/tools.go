package tools

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all eksctl tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register utility tools (refresh cache, list commands, etc.)
	RegisterUtilityTools(s)

	// Register all dynamic tools (auto-discovered from eksctl CLI)
	// This includes standard commands like create, get, delete, etc.
	err := RegisterAllDynamicTools(s)
	if err != nil {
		// Just log the error and continue
		// Some commands might fail to register but others will still work
		// This is better than failing completely
	}
}
