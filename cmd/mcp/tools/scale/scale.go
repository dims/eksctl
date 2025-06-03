package scale

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all scale tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register scale nodegroup command
	RegisterScaleNodegroupTool(s)
}
