package enable

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all enable tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register enable flux command
	RegisterEnableFluxTool(s)
}
