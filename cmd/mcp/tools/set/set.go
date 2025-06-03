package set

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all set tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register set labels command
	RegisterSetLabelsTool(s)
}
