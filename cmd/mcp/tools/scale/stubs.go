package scale

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

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
		cluster := request.GetString("cluster", "")
		region := request.GetString("region", "")
		name := request.GetString("name", "")
		
		if cluster == "" || region == "" || name == "" {
			return mcp.NewToolResultError("cluster, region, and name are required"), nil
		}
		
		return mcp.NewToolResultText(fmt.Sprintf("Scaling nodegroup %s in cluster %s in region %s (stub implementation)", name, cluster, region)), nil
	})
}
