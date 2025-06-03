package info

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers the info tool with the MCP server
func RegisterTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_info",
		mcp.WithDescription("Display information about eksctl"),
		mcp.WithString("output",
			mcp.Description("specifies the output format (valid option: json)"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Build the command arguments
		args := []string{"info"}

		// Add output format if specified
		output := request.GetString("output", "")
		if output != "" {
			args = append(args, "--output", output)
		}

		// Execute the eksctl command
		return common.ExecuteEksctlCommand(ctx, args)
	})
}
