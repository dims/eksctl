package register

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

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
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		provider := request.GetString("provider", "")
		
		if name == "" || region == "" || provider == "" {
			return mcp.NewToolResultError("name, region, and provider are required"), nil
		}
		
		return mcp.NewToolResultText(fmt.Sprintf("Registering %s cluster %s in region %s (stub implementation)", provider, name, region)), nil
	})
}
