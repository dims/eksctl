package deregister

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all deregister tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register deregister cluster command
	RegisterDeregisterClusterTool(s)
}
