package delete

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterDeleteClusterTool registers the eksctl_delete_cluster tool with the MCP server
func RegisterDeleteClusterTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_delete_cluster",
		mcp.WithDescription("Delete an Amazon EKS cluster and all associated resources"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name to delete"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("wait", 
			mcp.Description("Wait for cluster deletion to complete before returning (recommended)"),
		),
		mcp.WithString("force", 
			mcp.Description("Force deletion even if there are still cluster resources"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		
		params := map[string]string{
			"name":   name,
			"region": region,
		}
		
		wait := request.GetString("wait", "")
		if wait != "" {
			params["wait"] = wait
		}
		
		force := request.GetString("force", "")
		if force != "" {
			params["force"] = force
		}
		
		return common.CreateStubResponse(ctx, "delete cluster", params)
	})
}

// RegisterDeleteNodegroupTool registers the eksctl_delete_nodegroup tool with the MCP server
func RegisterDeleteNodegroupTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_delete_nodegroup",
		mcp.WithDescription("Delete a nodegroup from an Amazon EKS cluster"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name containing the nodegroup"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("nodegroup", 
			mcp.Description("Name of the nodegroup to delete"),
			mcp.Required(),
		),
		mcp.WithString("drain", 
			mcp.Description("Drain the nodegroup before deletion (evicts all pods)"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		nodegroup := request.GetString("nodegroup", "")
		
		params := map[string]string{
			"name":      name,
			"region":    region,
			"nodegroup": nodegroup,
		}
		
		drain := request.GetString("drain", "")
		if drain != "" {
			params["drain"] = drain
		}
		
		return common.CreateStubResponse(ctx, "delete nodegroup", params)
	})
}
