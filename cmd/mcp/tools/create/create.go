package create

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers all create tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register create cluster command
	RegisterCreateClusterTool(s)

	// Register create nodegroup command
	RegisterCreateNodegroupTool(s)

	// Register create iamserviceaccount command
	RegisterCreateIAMServiceAccountTool(s)

	// Register create iamidentitymapping command
	RegisterCreateIAMIdentityMappingTool(s)

	// Register create fargateprofile command
	RegisterCreateFargateProfileTool(s)

	// Register create addon command
	RegisterCreateAddonTool(s)

	// Register create accessentry command
	RegisterCreateAccessEntryTool(s)

	// Register create podidentityassociation command
	RegisterCreatePodIdentityAssociationTool(s)
}

// RegisterCreateClusterTool registers the eksctl_create_cluster tool with the MCP server
func RegisterCreateClusterTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_create_cluster",
		mcp.WithDescription("Create a new Amazon EKS cluster with default nodegroup"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name (must contain only alphanumeric characters and hyphens)"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster will be created"),
			mcp.Required(),
		),
		mcp.WithString("version", 
			mcp.Description("Kubernetes version (e.g. 1.28, 1.29, or 'latest')"),
		),
		mcp.WithString("nodegroup-name", 
			mcp.Description("Name for the initial managed nodegroup"),
		),
		mcp.WithNumber("nodes", 
			mcp.Description("Number of nodes in the initial nodegroup (default: 2)"),
		),
		mcp.WithNumber("nodes-min", 
			mcp.Description("Minimum number of nodes for auto-scaling"),
		),
		mcp.WithNumber("nodes-max", 
			mcp.Description("Maximum number of nodes for auto-scaling"),
		),
		mcp.WithString("node-type", 
			mcp.Description("EC2 instance type for the nodes (e.g. m5.large, t3.medium)"),
		),
		mcp.WithString("node-volume-size", 
			mcp.Description("Node volume size in GB (default: 80)"),
		),
		mcp.WithString("node-volume-type", 
			mcp.Description("Node volume type (gp2, gp3, io1, etc.)"),
		),
		mcp.WithString("ssh-access", 
			mcp.Description("Control SSH access for nodes (true or false)"),
		),
		mcp.WithString("ssh-public-key", 
			mcp.Description("SSH public key to use for nodes"),
		),
		mcp.WithString("vpc-cidr", 
			mcp.Description("Global CIDR to use for VPC"),
		),
		mcp.WithBoolean("without-nodegroup", 
			mcp.Description("If true, initial nodegroup will not be created"),
		),
		mcp.WithBoolean("dry-run", 
			mcp.Description("Do not apply any change, only show what would be done"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		version := request.GetString("version", "")
		
		params := map[string]string{
			"name":    name,
			"region":  region,
		}
		
		if version != "" {
			params["version"] = version
		}
		
		return common.CreateStubResponse(ctx, "create cluster", params)
	})
}

// RegisterCreateNodegroupTool registers the eksctl_create_nodegroup tool with the MCP server
func RegisterCreateNodegroupTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_create_nodegroup",
		mcp.WithDescription("Create a new nodegroup for an existing EKS cluster"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("nodegroup-name", 
			mcp.Description("Nodegroup name"),
			mcp.Required(),
		),
		mcp.WithNumber("nodes", 
			mcp.Description("Number of nodes"),
		),
		mcp.WithNumber("nodes-min", 
			mcp.Description("Minimum number of nodes"),
		),
		mcp.WithNumber("nodes-max", 
			mcp.Description("Maximum number of nodes"),
		),
		mcp.WithString("node-type", 
			mcp.Description("Node instance type"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		nodegroupName := request.GetString("nodegroup-name", "")
		
		params := map[string]string{
			"name":           name,
			"region":         region,
			"nodegroup-name": nodegroupName,
		}
		
		return common.CreateStubResponse(ctx, "create nodegroup", params)
	})
}

// RegisterCreateIAMServiceAccountTool registers the eksctl_create_iamserviceaccount tool with the MCP server
func RegisterCreateIAMServiceAccountTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_create_iamserviceaccount",
		mcp.WithDescription("Create an IAM service account"),
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
		
		return common.CreateStubResponse(ctx, "create iamserviceaccount", params)
	})
}

// RegisterCreateIAMIdentityMappingTool registers the eksctl_create_iamidentitymapping tool with the MCP server
func RegisterCreateIAMIdentityMappingTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_create_iamidentitymapping",
		mcp.WithDescription("Create an IAM identity mapping"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("arn", 
			mcp.Description("ARN of the IAM role or user to add"),
			mcp.Required(),
		),
		mcp.WithString("username", 
			mcp.Description("Kubernetes username"),
			mcp.Required(),
		),
		mcp.WithString("group", 
			mcp.Description("Kubernetes group (can be specified multiple times)"),
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
		
		group := request.GetString("group", "")
		if group != "" {
			params["group"] = group
		}
		
		return common.CreateStubResponse(ctx, "create iamidentitymapping", params)
	})
}

// RegisterCreateFargateProfileTool registers the eksctl_create_fargateprofile tool with the MCP server
func RegisterCreateFargateProfileTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_create_fargateprofile",
		mcp.WithDescription("Create a Fargate profile"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("profile-name", 
			mcp.Description("Fargate profile name"),
			mcp.Required(),
		),
		mcp.WithString("namespace", 
			mcp.Description("Kubernetes namespace"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		profileName := request.GetString("profile-name", "")
		namespace := request.GetString("namespace", "")
		
		params := map[string]string{
			"name":         name,
			"region":       region,
			"profile-name": profileName,
			"namespace":    namespace,
		}
		
		return common.CreateStubResponse(ctx, "create fargateprofile", params)
	})
}

// RegisterCreateAddonTool registers the eksctl_create_addon tool with the MCP server
func RegisterCreateAddonTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_create_addon",
		mcp.WithDescription("Create an addon"),
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
		mcp.WithString("version", 
			mcp.Description("Addon version"),
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
		
		version := request.GetString("version", "")
		if version != "" {
			params["version"] = version
		}
		
		return common.CreateStubResponse(ctx, "create addon", params)
	})
}

// RegisterCreateAccessEntryTool registers the eksctl_create_accessentry tool with the MCP server
func RegisterCreateAccessEntryTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_create_accessentry",
		mcp.WithDescription("Create an access entry"),
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
		mcp.WithString("kubernetes-groups", 
			mcp.Description("Kubernetes groups"),
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
		
		k8sGroups := request.GetString("kubernetes-groups", "")
		if k8sGroups != "" {
			params["kubernetes-groups"] = k8sGroups
		}
		
		return common.CreateStubResponse(ctx, "create accessentry", params)
	})
}

// RegisterCreatePodIdentityAssociationTool registers the eksctl_create_podidentityassociation tool with the MCP server
func RegisterCreatePodIdentityAssociationTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_create_podidentityassociation",
		mcp.WithDescription("Create a pod identity association"),
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
		mcp.WithString("role-arn", 
			mcp.Description("ARN of the IAM role"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		namespace := request.GetString("namespace", "")
		serviceAccount := request.GetString("service-account", "")
		roleARN := request.GetString("role-arn", "")
		
		params := map[string]string{
			"name":            name,
			"region":          region,
			"namespace":       namespace,
			"service-account": serviceAccount,
			"role-arn":        roleARN,
		}
		
		return common.CreateStubResponse(ctx, "create podidentityassociation", params)
	})
}
