package drain

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all drain tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register drain nodegroup command
	RegisterDrainNodegroupTool(s)
}

// RegisterDrainNodegroupTool registers the eksctl_drain_nodegroup tool with the MCP server
func RegisterDrainNodegroupTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_drain_nodegroup",
		mcp.WithDescription("Drain a nodegroup"),
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
		mcp.WithBoolean("disable-eviction",
			mcp.Description("Force drain without evicting pods"),
		),
		mcp.WithString("parallel",
			mcp.Description("Number of nodes to drain in parallel"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		cluster := request.GetString("cluster", "")
		region := request.GetString("region", "")
		name := request.GetString("name", "")

		if cluster == "" || region == "" || name == "" {
			return mcp.NewToolResultError("cluster, region, and name are required"), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Draining nodegroup %s in cluster %s in region %s (stub implementation)", name, cluster, region)), nil
	})
}
