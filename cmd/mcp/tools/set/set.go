package set

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers all set tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register set labels command
	RegisterSetLabelsTool(s)
}

// RegisterSetLabelsTool registers the eksctl_set_labels tool with the MCP server
func RegisterSetLabelsTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_set_labels",
		mcp.WithDescription("Set labels for a nodegroup"),
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
			mcp.Description("Labels to set (key=value,key2=value2)"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "set labels", request)
	})
}
