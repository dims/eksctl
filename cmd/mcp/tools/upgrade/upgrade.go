package upgrade

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all upgrade tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register upgrade cluster command
	RegisterUpgradeClusterTool(s)
	
	// Register upgrade nodegroup command
	RegisterUpgradeNodegroupTool(s)
}
