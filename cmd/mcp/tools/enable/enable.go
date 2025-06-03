package enable

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers all enable tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register enable flux command
	RegisterEnableFluxTool(s)
}

// RegisterEnableFluxTool registers the eksctl_enable_flux tool with the MCP server
func RegisterEnableFluxTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_enable_flux",
		mcp.WithDescription("Enable Flux for a cluster"),
		mcp.WithString("name",
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("git-url",
			mcp.Description("Git repository URL"),
			mcp.Required(),
		),
		mcp.WithString("git-branch",
			mcp.Description("Git branch"),
		),
		mcp.WithString("git-path",
			mcp.Description("Git path"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "enable flux", request)
	})
}
