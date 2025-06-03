package delete

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

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
		mcp.WithString("principal-arn", 
			mcp.Description("ARN of the principal"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		principalARN := request.GetString("principal-arn", "")
		
		params := map[string]string{
			"name":          name,
			"region":        region,
			"principal-arn": principalARN,
		}
		
		return common.CreateStubResponse(ctx, "delete accessentry", params)
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
		
		return common.CreateStubResponse(ctx, "delete podidentityassociation", params)
	})
}
