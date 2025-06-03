package completion

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterCompletionBashTool registers the eksctl_completion_bash tool with the MCP server
func RegisterCompletionBashTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_completion_bash",
		mcp.WithDescription("Generate bash completion"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("# Bash completion script for eksctl (stub implementation)"), nil
	})
}

// RegisterCompletionFishTool registers the eksctl_completion_fish tool with the MCP server
func RegisterCompletionFishTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_completion_fish",
		mcp.WithDescription("Generate fish completion"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("# Fish completion script for eksctl (stub implementation)"), nil
	})
}

// RegisterCompletionPowershellTool registers the eksctl_completion_powershell tool with the MCP server
func RegisterCompletionPowershellTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_completion_powershell",
		mcp.WithDescription("Generate powershell completion"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("# PowerShell completion script for eksctl (stub implementation)"), nil
	})
}

// RegisterCompletionZshTool registers the eksctl_completion_zsh tool with the MCP server
func RegisterCompletionZshTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_completion_zsh",
		mcp.WithDescription("Generate zsh completion"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("# Zsh completion script for eksctl (stub implementation)"), nil
	})
}
