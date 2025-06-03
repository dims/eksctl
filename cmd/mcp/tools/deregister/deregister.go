package deregister

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers all deregister tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register deregister cluster command
	RegisterDeregisterClusterTool(s)
}

// RegisterDeregisterClusterTool registers the eksctl_deregister_cluster tool with the MCP server
func RegisterDeregisterClusterTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_deregister_cluster",
		mcp.WithDescription("Deregister a non-EKS cluster"),
		mcp.WithString("name",
			mcp.Description("Cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "deregister cluster", request)
	})
}
