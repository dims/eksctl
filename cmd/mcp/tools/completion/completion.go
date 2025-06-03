package completion

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all completion tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register completion bash command
	RegisterCompletionBashTool(s)
	
	// Register completion fish command
	RegisterCompletionFishTool(s)
	
	// Register completion powershell command
	RegisterCompletionPowershellTool(s)
	
	// Register completion zsh command
	RegisterCompletionZshTool(s)
}
