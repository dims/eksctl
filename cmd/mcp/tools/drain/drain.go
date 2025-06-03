package drain

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all drain tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register drain nodegroup command
	RegisterDrainNodegroupTool(s)
}
