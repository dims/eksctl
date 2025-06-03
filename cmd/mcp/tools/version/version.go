package version

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/pkg/version"
)

// RegisterTools registers the version tool with the MCP server
func RegisterTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_version",
		mcp.WithDescription("Output the version of eksctl"),
		mcp.WithString("output",
			mcp.Description("specifies the output format (valid option: json)"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		output := request.GetString("output", "")

		switch output {
		case "":
			return mcp.NewToolResultText(version.GetVersion()), nil
		case "json":
			return mcp.NewToolResultText(version.String()), nil
		default:
			return mcp.NewToolResultError("unknown output: " + output), nil
		}
	})
}
