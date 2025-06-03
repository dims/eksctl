package get

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterGetClusterTool registers the eksctl_get_cluster tool with the MCP server
func RegisterGetClusterTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_cluster",
		mcp.WithDescription("Get EKS cluster(s)"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		
		if region == "" {
			return mcp.NewToolResultError("region is required"), nil
		}
		
		if name != "" {
			return mcp.NewToolResultText(fmt.Sprintf("Cluster %s in region %s (stub implementation)", name, region)), nil
		} else {
			return mcp.NewToolResultText(fmt.Sprintf("All clusters in region %s (stub implementation)", region)), nil
		}
	})
}

// RegisterGetNodegroupTool registers the eksctl_get_nodegroup tool with the MCP server
func RegisterGetNodegroupTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_nodegroup",
		mcp.WithDescription("Get nodegroup(s)"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		
		if name == "" || region == "" {
			return mcp.NewToolResultError("name and region are required"), nil
		}
		
		return mcp.NewToolResultText(fmt.Sprintf("Nodegroups for cluster %s in region %s (stub implementation)", name, region)), nil
	})
}

// RegisterGetIAMServiceAccountTool registers the eksctl_get_iamserviceaccount tool with the MCP server
func RegisterGetIAMServiceAccountTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_iamserviceaccount",
		mcp.WithDescription("Get IAM service account(s)"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("This command is not yet implemented"), nil
	})
}

// RegisterGetIAMIdentityMappingTool registers the eksctl_get_iamidentitymapping tool with the MCP server
func RegisterGetIAMIdentityMappingTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_iamidentitymapping",
		mcp.WithDescription("Get IAM identity mapping(s)"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("This command is not yet implemented"), nil
	})
}

// RegisterGetFargateProfileTool registers the eksctl_get_fargateprofile tool with the MCP server
func RegisterGetFargateProfileTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_fargateprofile",
		mcp.WithDescription("Get Fargate profile(s)"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("This command is not yet implemented"), nil
	})
}

// RegisterGetAddonTool registers the eksctl_get_addon tool with the MCP server
func RegisterGetAddonTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_addon",
		mcp.WithDescription("Get addon(s)"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("This command is not yet implemented"), nil
	})
}

// RegisterGetAccessEntryTool registers the eksctl_get_accessentry tool with the MCP server
func RegisterGetAccessEntryTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_accessentry",
		mcp.WithDescription("Get access entry(s)"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("This command is not yet implemented"), nil
	})
}

// RegisterGetPodIdentityAssociationTool registers the eksctl_get_podidentityassociation tool with the MCP server
func RegisterGetPodIdentityAssociationTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_podidentityassociation",
		mcp.WithDescription("Get pod identity association(s)"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("This command is not yet implemented"), nil
	})
}
