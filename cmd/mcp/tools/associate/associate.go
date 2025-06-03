package associate

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all associate tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register associate identityprovider command
	RegisterAssociateIdentityProviderTool(s)
}
