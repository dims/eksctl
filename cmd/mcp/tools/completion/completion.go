package completion

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
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

// RegisterCompletionBashTool registers the eksctl_completion_bash tool with the MCP server
func RegisterCompletionBashTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_completion_bash",
		mcp.WithDescription("Generate bash completion"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommand(ctx, []string{"completion", "bash"})
	})
}

// RegisterCompletionFishTool registers the eksctl_completion_fish tool with the MCP server
func RegisterCompletionFishTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_completion_fish",
		mcp.WithDescription("Generate fish completion"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommand(ctx, []string{"completion", "fish"})
	})
}

// RegisterCompletionPowershellTool registers the eksctl_completion_powershell tool with the MCP server
func RegisterCompletionPowershellTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_completion_powershell",
		mcp.WithDescription("Generate powershell completion"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommand(ctx, []string{"completion", "powershell"})
	})
}

// RegisterCompletionZshTool registers the eksctl_completion_zsh tool with the MCP server
func RegisterCompletionZshTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_completion_zsh",
		mcp.WithDescription("Generate zsh completion"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommand(ctx, []string{"completion", "zsh"})
	})
}
