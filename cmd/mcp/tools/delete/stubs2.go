package delete

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

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
		mcp.WithString("namespace", 
			mcp.Description("Kubernetes namespace"),
			mcp.Required(),
		),
		mcp.WithString("service-account", 
			mcp.Description("Kubernetes service account name"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		namespace := request.GetString("namespace", "")
		serviceAccount := request.GetString("service-account", "")
		
		params := map[string]string{
			"name":            name,
			"region":          region,
			"namespace":       namespace,
			"service-account": serviceAccount,
		}
		
		return common.CreateStubResponse(ctx, "delete iamserviceaccount", params)
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
		mcp.WithString("arn", 
			mcp.Description("ARN of the IAM role or user to remove"),
			mcp.Required(),
		),
		mcp.WithString("username", 
			mcp.Description("Kubernetes username"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		arn := request.GetString("arn", "")
		username := request.GetString("username", "")
		
		params := map[string]string{
			"name":     name,
			"region":   region,
			"arn":      arn,
			"username": username,
		}
		
		return common.CreateStubResponse(ctx, "delete iamidentitymapping", params)
	})
}
