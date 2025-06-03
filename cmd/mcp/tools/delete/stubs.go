package delete

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterDeleteClusterTool registers the eksctl_delete_cluster tool with the MCP server
func RegisterDeleteClusterTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_delete_cluster",
		mcp.WithDescription("Delete an EKS cluster"),
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
		
		return mcp.NewToolResultText(fmt.Sprintf("Deleting cluster %s in region %s (stub implementation)", name, region)), nil
	})
}

// RegisterDeleteNodegroupTool registers the eksctl_delete_nodegroup tool with the MCP server
func RegisterDeleteNodegroupTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_delete_nodegroup",
		mcp.WithDescription("Delete a nodegroup from an EKS cluster"),
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
		
		return mcp.NewToolResultText(fmt.Sprintf("Deleting nodegroup from cluster %s in region %s (stub implementation)", name, region)), nil
	})
}

// RegisterDeleteIAMServiceAccountTool registers the eksctl_delete_iamserviceaccount tool with the MCP server
func RegisterDeleteIAMServiceAccountTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_delete_iamserviceaccount",
		mcp.WithDescription("Delete an IAM service account"),
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

// RegisterDeleteIAMIdentityMappingTool registers the eksctl_delete_iamidentitymapping tool with the MCP server
func RegisterDeleteIAMIdentityMappingTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_delete_iamidentitymapping",
		mcp.WithDescription("Delete an IAM identity mapping"),
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
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("This command is not yet implemented"), nil
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
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return mcp.NewToolResultText("This command is not yet implemented"), nil
	})
}

// RegisterDeleteAccessEntryTool registers the eksctl_delete_accessentry tool with the MCP server
func RegisterDeleteAccessEntryTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_delete_accessentry",
		mcp.WithDescription("Delete an access entry"),
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

// RegisterDeletePodIdentityAssociationTool registers the eksctl_delete_podidentityassociation tool with the MCP server
func RegisterDeletePodIdentityAssociationTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_delete_podidentityassociation",
		mcp.WithDescription("Delete a pod identity association"),
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
