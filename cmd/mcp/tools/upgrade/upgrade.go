package upgrade

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers all upgrade tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register upgrade cluster command
	RegisterUpgradeClusterTool(s)

	// Register upgrade nodegroup command
	RegisterUpgradeNodegroupTool(s)
}

// RegisterUpgradeClusterTool registers the eksctl_upgrade_cluster tool with the MCP server
func RegisterUpgradeClusterTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_upgrade_cluster",
		mcp.WithDescription("Upgrade a cluster control plane to a new Kubernetes version"),
		mcp.WithString("name",
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("version",
			mcp.Description("Kubernetes version"),
			mcp.Required(),
		),
		mcp.WithBoolean("approve",
			mcp.Description("Skip confirmation prompt"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "upgrade cluster", request)
	})
}

// RegisterUpgradeNodegroupTool registers the eksctl_upgrade_nodegroup tool with the MCP server
func RegisterUpgradeNodegroupTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_upgrade_nodegroup",
		mcp.WithDescription("Upgrade a nodegroup to match the control plane Kubernetes version"),
		mcp.WithString("name",
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("nodegroup",
			mcp.Description("Nodegroup name"),
			mcp.Required(),
		),
		mcp.WithString("kubernetes-version",
			mcp.Description("Kubernetes version"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "upgrade nodegroup", request)
	})
}
