package unset

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)



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
		cluster := request.GetString("cluster", "")
		region := request.GetString("region", "")
		nodegroup := request.GetString("nodegroup", "")
		labels := request.GetString("labels", "")
		
		if cluster == "" || region == "" || nodegroup == "" || labels == "" {
			return mcp.NewToolResultError("cluster, region, nodegroup, and labels are required"), nil
		}
		
		return mcp.NewToolResultText(fmt.Sprintf("Unsetting labels %s for nodegroup %s in cluster %s in region %s (stub implementation)", labels, nodegroup, cluster, region)), nil
	})
}


// RegisterTools registers all unset tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register unset labels command
	RegisterUnsetLabelsTool(s)
}
