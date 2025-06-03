package delete

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterDeleteFargateProfileTool registers the eksctl_delete_fargateprofile tool with the MCP server
func RegisterDeleteFargateProfileTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_delete_fargateprofile",
		mcp.WithDescription("Delete a Fargate profile"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("profile", 
			mcp.Description("Fargate profile name"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		profile := request.GetString("profile", "")
		
		params := map[string]string{
			"name":    name,
			"region":  region,
			"profile": profile,
		}
		
		return common.CreateStubResponse(ctx, "delete fargateprofile", params)
	})
}

// RegisterDeleteAddonTool registers the eksctl_delete_addon tool with the MCP server
func RegisterDeleteAddonTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_delete_addon",
		mcp.WithDescription("Delete an addon"),
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
		mcp.WithString("preserve", 
			mcp.Description("Preserve the resources created by the addon"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		addon := request.GetString("addon", "")
		
		params := map[string]string{
			"name":   name,
			"region": region,
			"addon":  addon,
		}
		
		preserve := request.GetString("preserve", "")
		if preserve != "" {
			params["preserve"] = preserve
		}
		
		return common.CreateStubResponse(ctx, "delete addon", params)
	})
}
