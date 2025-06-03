package register

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers all register tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register register cluster command
	RegisterRegisterClusterTool(s)
}

// RegisterRegisterClusterTool registers the eksctl_register_cluster tool with the MCP server
func RegisterRegisterClusterTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_register_cluster",
		mcp.WithDescription("Register a non-EKS cluster"),
		mcp.WithString("name",
			mcp.Description("Cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("provider",
			mcp.Description("Kubernetes provider"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "register cluster", request)
	})
}
