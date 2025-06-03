package update

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers all update tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register update addon command
	RegisterUpdateAddonTool(s)

	// Register update cluster command
	RegisterUpdateClusterTool(s)

	// Register update cluster-logging command
	RegisterUpdateClusterLoggingTool(s)

	// Register update cluster-config command
	RegisterUpdateClusterConfigTool(s)

	// Register update nodegroup command
	RegisterUpdateNodegroupTool(s)

	// Register update accessentry command
	RegisterUpdateAccessEntryTool(s)
}

// RegisterUpdateAddonTool registers the eksctl_update_addon tool with the MCP server
func RegisterUpdateAddonTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_update_addon",
		mcp.WithDescription("Update an addon"),
		mcp.WithString("name",
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("addon",
			mcp.Description("Addon name"),
			mcp.Required(),
		),
		mcp.WithString("version",
			mcp.Description("Addon version"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "update addon", request)
	})
}

// RegisterUpdateClusterTool registers the eksctl_update_cluster tool with the MCP server
func RegisterUpdateClusterTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_update_cluster",
		mcp.WithDescription("Update cluster configuration"),
		mcp.WithString("name",
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("timeout",
			mcp.Description("Maximum waiting time for operations"),
		),
		mcp.WithString("approve",
			mcp.Description("Skip confirmation prompt"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "update cluster", request)
	})
}

// RegisterUpdateClusterLoggingTool registers the eksctl_update_cluster-logging tool with the MCP server
func RegisterUpdateClusterLoggingTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_update_cluster-logging",
		mcp.WithDescription("Update cluster logging configuration"),
		mcp.WithString("name",
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("enable-types",
			mcp.Description("Enable logging for types (api,audit,authenticator,controllerManager,scheduler)"),
		),
		mcp.WithString("disable-types",
			mcp.Description("Disable logging for types (api,audit,authenticator,controllerManager,scheduler)"),
		),
		mcp.WithString("approve",
			mcp.Description("Skip confirmation prompt"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "update cluster-logging", request)
	})
}

// RegisterUpdateClusterConfigTool registers the eksctl_update_cluster-config tool with the MCP server
func RegisterUpdateClusterConfigTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_update_cluster-config",
		mcp.WithDescription("Update cluster configuration"),
		mcp.WithString("name",
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("config-file",
			mcp.Description("Path to the cluster config file"),
		),
		mcp.WithString("approve",
			mcp.Description("Skip confirmation prompt"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "update cluster-config", request)
	})
}

// RegisterUpdateNodegroupTool registers the eksctl_update_nodegroup tool with the MCP server
func RegisterUpdateNodegroupTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_update_nodegroup",
		mcp.WithDescription("Update nodegroup configuration"),
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
		mcp.WithString("max-pods-per-node",
			mcp.Description("Maximum number of pods per node"),
		),
		mcp.WithString("timeout",
			mcp.Description("Maximum waiting time for operations"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "update nodegroup", request)
	})
}

// RegisterUpdateAccessEntryTool registers the eksctl_update_accessentry tool with the MCP server
func RegisterUpdateAccessEntryTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_update_accessentry",
		mcp.WithDescription("Update an access entry"),
		mcp.WithString("name",
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("principal-arn",
			mcp.Description("ARN of the principal"),
			mcp.Required(),
		),
		mcp.WithString("kubernetes-groups",
			mcp.Description("Kubernetes groups"),
		),
		mcp.WithString("type",
			mcp.Description("Type of access entry (STANDARD or CLUSTER_ADMIN)"),
		),
		mcp.WithString("access-policy-arn",
			mcp.Description("ARN of the access policy"),
		),
		mcp.WithString("access-scope",
			mcp.Description("Access scope (namespace or cluster)"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "update accessentry", request)
	})
}
