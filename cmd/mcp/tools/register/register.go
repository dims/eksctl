package register

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all register tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register register cluster command
	RegisterRegisterClusterTool(s)
}
