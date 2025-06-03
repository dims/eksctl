package disassociate

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all disassociate tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register disassociate identityprovider command
	RegisterDisassociateIdentityProviderTool(s)
}
