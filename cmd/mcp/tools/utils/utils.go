package utils

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all utils tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register utils write-kubeconfig command
	RegisterWriteKubeconfigTool(s)

	// Register utils describe-stacks command
	RegisterDescribeStacksTool(s)

	// Register utils update-kube-proxy command
	RegisterUpdateKubeProxyTool(s)

	// Register utils update-aws-node command
	RegisterUpdateAWSNodeTool(s)

	// Register utils update-coredns command
	RegisterUpdateCoreDNSTool(s)

	// Register utils associate-iam-oidc-provider command
	RegisterAssociateIAMOIDCProviderTool(s)
}

// RegisterWriteKubeconfigTool registers the eksctl_utils_write-kubeconfig tool with the MCP server
func RegisterWriteKubeconfigTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_utils_write-kubeconfig",
		mcp.WithDescription("Write kubeconfig file for an Amazon EKS cluster to enable kubectl access"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name to generate kubeconfig for"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("kubeconfig", 
			mcp.Description("Path to write kubeconfig file (default: ~/.kube/config)"),
		),
		mcp.WithString("authenticator-role-arn", 
			mcp.Description("AWS IAM role to assume for authenticator"),
		),
		mcp.WithString("set-kubeconfig-context", 
			mcp.Description("Set current-context in kubeconfig (default: true)"),
		),
		mcp.WithString("auto-kubeconfig", 
			mcp.Description("Save kubeconfig file by cluster name"),
		),
		mcp.WithString("write-kubeconfig", 
			mcp.Description("Toggle writing of kubeconfig (default: true)"),
		),
		mcp.WithString("profile", 
			mcp.Description("AWS credentials profile to use"),
		),
		mcp.WithString("timeout", 
			mcp.Description("Maximum waiting time for operations"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		
		if name == "" || region == "" {
			return mcp.NewToolResultError("name and region are required"), nil
		}
		
		return mcp.NewToolResultText(fmt.Sprintf("Writing kubeconfig for cluster %s in region %s (stub implementation)", name, region)), nil
	})
}

// RegisterDescribeStacksTool registers the eksctl_utils_describe-stacks tool with the MCP server
func RegisterDescribeStacksTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_utils_describe-stacks",
		mcp.WithDescription("Describe CloudFormation stacks for an Amazon EKS cluster"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("output", 
			mcp.Description("Output format: json or yaml"),
		),
		mcp.WithString("profile", 
			mcp.Description("AWS credentials profile to use"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		
		if name == "" || region == "" {
			return mcp.NewToolResultError("name and region are required"), nil
		}
		
		return mcp.NewToolResultText(fmt.Sprintf("Describing CloudFormation stacks for cluster %s in region %s (stub implementation)", name, region)), nil
	})
}

// RegisterUpdateKubeProxyTool registers the eksctl_utils_update-kube-proxy tool with the MCP server
func RegisterUpdateKubeProxyTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_utils_update-kube-proxy",
		mcp.WithDescription("Update kube-proxy add-on to the latest version"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("approve", 
			mcp.Description("Skip confirmation prompt"),
		),
		mcp.WithString("profile", 
			mcp.Description("AWS credentials profile to use"),
		),
		mcp.WithString("timeout", 
			mcp.Description("Maximum waiting time for operations"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		
		if name == "" || region == "" {
			return mcp.NewToolResultError("name and region are required"), nil
		}
		
		return mcp.NewToolResultText(fmt.Sprintf("Updating kube-proxy add-on for cluster %s in region %s (stub implementation)", name, region)), nil
	})
}

// RegisterUpdateAWSNodeTool registers the eksctl_utils_update-aws-node tool with the MCP server
func RegisterUpdateAWSNodeTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_utils_update-aws-node",
		mcp.WithDescription("Update aws-node (Amazon VPC CNI) add-on to the latest version"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("approve", 
			mcp.Description("Skip confirmation prompt"),
		),
		mcp.WithString("profile", 
			mcp.Description("AWS credentials profile to use"),
		),
		mcp.WithString("timeout", 
			mcp.Description("Maximum waiting time for operations"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		
		if name == "" || region == "" {
			return mcp.NewToolResultError("name and region are required"), nil
		}
		
		return mcp.NewToolResultText(fmt.Sprintf("Updating aws-node add-on for cluster %s in region %s (stub implementation)", name, region)), nil
	})
}

// RegisterUpdateCoreDNSTool registers the eksctl_utils_update-coredns tool with the MCP server
func RegisterUpdateCoreDNSTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_utils_update-coredns",
		mcp.WithDescription("Update CoreDNS add-on to the latest version"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("approve", 
			mcp.Description("Skip confirmation prompt"),
		),
		mcp.WithString("profile", 
			mcp.Description("AWS credentials profile to use"),
		),
		mcp.WithString("timeout", 
			mcp.Description("Maximum waiting time for operations"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		
		if name == "" || region == "" {
			return mcp.NewToolResultError("name and region are required"), nil
		}
		
		return mcp.NewToolResultText(fmt.Sprintf("Updating CoreDNS add-on for cluster %s in region %s (stub implementation)", name, region)), nil
	})
}

// RegisterAssociateIAMOIDCProviderTool registers the eksctl_utils_associate-iam-oidc-provider tool with the MCP server
func RegisterAssociateIAMOIDCProviderTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_utils_associate-iam-oidc-provider",
		mcp.WithDescription("Associate IAM OIDC provider for an Amazon EKS cluster to enable IAM roles for service accounts"),
		mcp.WithString("name", 
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region", 
			mcp.Description("AWS region where the cluster is located"),
			mcp.Required(),
		),
		mcp.WithString("approve", 
			mcp.Description("Skip confirmation prompt"),
		),
		mcp.WithString("profile", 
			mcp.Description("AWS credentials profile to use"),
		),
		mcp.WithString("timeout", 
			mcp.Description("Maximum waiting time for operations"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		
		if name == "" || region == "" {
			return mcp.NewToolResultError("name and region are required"), nil
		}
		
		return mcp.NewToolResultText(fmt.Sprintf("Associating IAM OIDC provider for cluster %s in region %s (stub implementation)", name, region)), nil
	})
}
