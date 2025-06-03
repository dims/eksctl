package enable

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterEnableFluxTool registers the eksctl_enable_flux tool with the MCP server
func RegisterEnableFluxTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_enable_flux",
		mcp.WithDescription("Enable Flux for a cluster"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("git-url", 
			mcp.Description("Git repository URL"),
			mcp.Required(),
		),
		mcp.WithString("git-branch", 
			mcp.Description("Git branch"),
		),
		mcp.WithString("git-path", 
			mcp.Description("Git path"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		gitUrl := request.GetString("git-url", "")
		
		if name == "" || region == "" || gitUrl == "" {
			return mcp.NewToolResultError("name, region, and git-url are required"), nil
		}
		
		return mcp.NewToolResultText(fmt.Sprintf("Enabling Flux for cluster %s in region %s with git repository %s (stub implementation)", name, region, gitUrl)), nil
	})
}
