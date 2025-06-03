package help

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers the help tool with the MCP server
func RegisterTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_help",
		mcp.WithDescription("Help about any command"),
		mcp.WithString("command",
			mcp.Description("Command to get help for"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		command := request.GetString("command", "")

		// Build the command arguments
		args := []string{"help"}
		if command != "" {
			args = append(args, command)
		}

		// Execute the eksctl command
		return common.ExecuteEksctlCommand(ctx, args)
	})
}
