package unset

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all unset tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register unset labels command
	RegisterUnsetLabelsTool(s)
}
