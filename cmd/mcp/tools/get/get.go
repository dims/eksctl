package get

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)



// RegisterTools registers all get tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register get cluster command
	RegisterGetClusterTool(s)

	// Register get nodegroup command
	RegisterGetNodegroupTool(s)

	// Register get iamserviceaccount command
	RegisterGetIAMServiceAccountTool(s)

	// Register get iamidentitymapping command
	RegisterGetIAMIdentityMappingTool(s)

	// Register get fargateprofile command
	RegisterGetFargateProfileTool(s)

	// Register get addon command
	RegisterGetAddonTool(s)

	// Register get accessentry command
	RegisterGetAccessEntryTool(s)

	// Register get podidentityassociation command
	RegisterGetPodIdentityAssociationTool(s)
}


// RegisterGetClusterTool registers the eksctl_get_cluster tool with the MCP server
func RegisterGetClusterTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_cluster",
		mcp.WithDescription("Get details about Amazon EKS clusters"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name to get details for (if omitted, lists all clusters)"),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region to query"),
			mcp.Required(),
		),
		mcp.WithString("output", 
			mcp.Description("Output format: table, json, or yaml (default: table)"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		output := request.GetString("output", "")
		
		params := map[string]string{
			"region": region,
		}
		
		if name != "" {
			params["name"] = name
		}
		
		if output != "" {
			params["output"] = output
		}
		
		return common.CreateStubResponse(ctx, "get cluster", params)
	})
}

// RegisterGetNodegroupTool registers the eksctl_get_nodegroup tool with the MCP server
func RegisterGetNodegroupTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_nodegroup",
		mcp.WithDescription("Get details about nodegroups in an Amazon EKS cluster"),
		mcp.WithString("name", 
			mcp.Description("Nodegroup name to get details for (if omitted, lists all nodegroups)"),
		),
		mcp.WithString("cluster", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("output", 
			mcp.Description("Output format: table, json, or yaml (default: table)"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		cluster := request.GetString("cluster", "")
		region := request.GetString("region", "")
		output := request.GetString("output", "")
		
		params := map[string]string{
			"cluster": cluster,
			"region":  region,
		}
		
		if name != "" {
			params["name"] = name
		}
		
		if output != "" {
			params["output"] = output
		}
		
		return common.CreateStubResponse(ctx, "get nodegroup", params)
	})
}

// RegisterGetIAMServiceAccountTool registers the eksctl_get_iamserviceaccount tool with the MCP server
func RegisterGetIAMServiceAccountTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_iamserviceaccount",
		mcp.WithDescription("Get IAM service accounts for an Amazon EKS cluster"),
		mcp.WithString("name", 
			mcp.Description("Service account name to get details for (if omitted, lists all service accounts)"),
		),
		mcp.WithString("cluster", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("namespace", 
			mcp.Description("Kubernetes namespace to filter by"),
		),
		mcp.WithString("output", 
			mcp.Description("Output format: table, json, or yaml (default: table)"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		cluster := request.GetString("cluster", "")
		region := request.GetString("region", "")
		namespace := request.GetString("namespace", "")
		output := request.GetString("output", "")
		
		params := map[string]string{
			"cluster": cluster,
			"region":  region,
		}
		
		if name != "" {
			params["name"] = name
		}
		
		if namespace != "" {
			params["namespace"] = namespace
		}
		
		if output != "" {
			params["output"] = output
		}
		
		return common.CreateStubResponse(ctx, "get iamserviceaccount", params)
	})
}

// RegisterGetIAMIdentityMappingTool registers the eksctl_get_iamidentitymapping tool with the MCP server
func RegisterGetIAMIdentityMappingTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_iamidentitymapping",
		mcp.WithDescription("Get IAM identity mappings for an Amazon EKS cluster"),
		mcp.WithString("cluster", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("arn", 
			mcp.Description("ARN of the IAM role or user to get mappings for"),
		),
		mcp.WithString("output", 
			mcp.Description("Output format: table, json, or yaml (default: table)"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		cluster := request.GetString("cluster", "")
		region := request.GetString("region", "")
		arn := request.GetString("arn", "")
		output := request.GetString("output", "")
		
		params := map[string]string{
			"cluster": cluster,
			"region":  region,
		}
		
		if arn != "" {
			params["arn"] = arn
		}
		
		if output != "" {
			params["output"] = output
		}
		
		return common.CreateStubResponse(ctx, "get iamidentitymapping", params)
	})
}

// RegisterGetFargateProfileTool registers the eksctl_get_fargateprofile tool with the MCP server
func RegisterGetFargateProfileTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_fargateprofile",
		mcp.WithDescription("Get Fargate profiles for an Amazon EKS cluster"),
		mcp.WithString("name", 
			mcp.Description("Fargate profile name to get details for (if omitted, lists all profiles)"),
		),
		mcp.WithString("cluster", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("output", 
			mcp.Description("Output format: table, json, or yaml (default: table)"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		cluster := request.GetString("cluster", "")
		region := request.GetString("region", "")
		output := request.GetString("output", "")
		
		params := map[string]string{
			"cluster": cluster,
			"region":  region,
		}
		
		if name != "" {
			params["name"] = name
		}
		
		if output != "" {
			params["output"] = output
		}
		
		return common.CreateStubResponse(ctx, "get fargateprofile", params)
	})
}

// RegisterGetAddonTool registers the eksctl_get_addon tool with the MCP server
func RegisterGetAddonTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_addon",
		mcp.WithDescription("Get EKS add-ons for an Amazon EKS cluster"),
		mcp.WithString("name", 
			mcp.Description("Add-on name to get details for (if omitted, lists all add-ons)"),
		),
		mcp.WithString("cluster", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("output", 
			mcp.Description("Output format: table, json, or yaml (default: table)"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		cluster := request.GetString("cluster", "")
		region := request.GetString("region", "")
		output := request.GetString("output", "")
		
		params := map[string]string{
			"cluster": cluster,
			"region":  region,
		}
		
		if name != "" {
			params["name"] = name
		}
		
		if output != "" {
			params["output"] = output
		}
		
		return common.CreateStubResponse(ctx, "get addon", params)
	})
}

// RegisterGetAccessEntryTool registers the eksctl_get_accessentry tool with the MCP server
func RegisterGetAccessEntryTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_accessentry",
		mcp.WithDescription("Get access entries for an Amazon EKS cluster"),
		mcp.WithString("cluster", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("principal-arn", 
			mcp.Description("ARN of the principal to get access entries for"),
		),
		mcp.WithString("output", 
			mcp.Description("Output format: table, json, or yaml (default: table)"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		cluster := request.GetString("cluster", "")
		region := request.GetString("region", "")
		principalARN := request.GetString("principal-arn", "")
		output := request.GetString("output", "")
		
		params := map[string]string{
			"cluster": cluster,
			"region":  region,
		}
		
		if principalARN != "" {
			params["principal-arn"] = principalARN
		}
		
		if output != "" {
			params["output"] = output
		}
		
		return common.CreateStubResponse(ctx, "get accessentry", params)
	})
}

// RegisterGetPodIdentityAssociationTool registers the eksctl_get_podidentityassociation tool with the MCP server
func RegisterGetPodIdentityAssociationTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_get_podidentityassociation",
		mcp.WithDescription("Get pod identity associations for an Amazon EKS cluster"),
		mcp.WithString("cluster", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("namespace", 
			mcp.Description("Kubernetes namespace to filter by"),
		),
		mcp.WithString("service-account", 
			mcp.Description("Kubernetes service account name to filter by"),
		),
		mcp.WithString("output", 
			mcp.Description("Output format: table, json, or yaml (default: table)"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		cluster := request.GetString("cluster", "")
		region := request.GetString("region", "")
		namespace := request.GetString("namespace", "")
		serviceAccount := request.GetString("service-account", "")
		output := request.GetString("output", "")
		
		params := map[string]string{
			"cluster": cluster,
			"region":  region,
		}
		
		if namespace != "" {
			params["namespace"] = namespace
		}
		
		if serviceAccount != "" {
			params["service-account"] = serviceAccount
		}
		
		if output != "" {
			params["output"] = output
		}
		
		return common.CreateStubResponse(ctx, "get podidentityassociation", params)
	})
}
