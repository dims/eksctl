package delete

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers all delete tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register delete cluster command
	RegisterDeleteClusterTool(s)

	// Register delete nodegroup command
	RegisterDeleteNodegroupTool(s)

	// Register delete iamserviceaccount command
	RegisterDeleteIAMServiceAccountTool(s)

	// Register delete iamidentitymapping command
	RegisterDeleteIAMIdentityMappingTool(s)

	// Register delete fargateprofile command
	RegisterDeleteFargateProfileTool(s)

	// Register delete addon command
	RegisterDeleteAddonTool(s)

	// Register delete accessentry command
	RegisterDeleteAccessEntryTool(s)

	// Register delete podidentityassociation command
	RegisterDeletePodIdentityAssociationTool(s)
}

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
