package tools

import (
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/dynamic"
)

// RegisterTools registers all eksctl tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register dynamic tools (auto-discovered from eksctl CLI)
	// This includes the anywhere command with special handling
	dynamic.RegisterTools(s)
}
