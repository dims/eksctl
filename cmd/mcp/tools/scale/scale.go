package scale

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers all scale tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register scale nodegroup command
	RegisterScaleNodegroupTool(s)
}

// RegisterScaleNodegroupTool registers the eksctl_scale_nodegroup tool with the MCP server
func RegisterScaleNodegroupTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_scale_nodegroup",
		mcp.WithDescription("Scale a nodegroup"),
		mcp.WithString("cluster",
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("Nodegroup name"),
			mcp.Required(),
		),
		mcp.WithNumber("nodes",
			mcp.Description("Desired number of nodes"),
		),
		mcp.WithNumber("nodes-min",
			mcp.Description("Minimum number of nodes"),
		),
		mcp.WithNumber("nodes-max",
			mcp.Description("Maximum number of nodes"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "scale nodegroup", request)
	})
}
