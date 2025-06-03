package unset

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers all unset tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register unset labels command
	RegisterUnsetLabelsTool(s)
}

// RegisterUnsetLabelsTool registers the eksctl_unset_labels tool with the MCP server
func RegisterUnsetLabelsTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_unset_labels",
		mcp.WithDescription("Unset labels for a nodegroup"),
		mcp.WithString("cluster",
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
		mcp.WithString("labels",
			mcp.Description("Labels to unset (key1,key2,key3)"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "unset labels", request)
	})
}
