package common

import (
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/server"
)

// DynamicToolCategory represents a category of related eksctl commands
type DynamicToolCategory struct {
	Name        string
	Description string
	Commands    []DynamicToolCommand
}

// DynamicToolCommand represents an eksctl command to be exposed as an MCP tool
type DynamicToolCommand struct {
	Name        string
	Description string
	Command     string
}

// RegisterDynamicToolCategory registers all commands in a category
func RegisterDynamicToolCategory(s *server.MCPServer, category DynamicToolCategory) error {
	for _, cmd := range category.Commands {
		toolName := fmt.Sprintf("eksctl_%s", cmd.Name)
		err := RegisterDynamicTool(s, toolName, cmd.Description, cmd.Command)
		if err != nil {
			return fmt.Errorf("failed to register tool %s: %w", toolName, err)
		}
	}
	return nil
}

// RegisterAllDynamicTools registers all eksctl commands as MCP tools
func RegisterAllDynamicTools(s *server.MCPServer) error {
	// Define all command categories
	categories := []DynamicToolCategory{
		{
			Name:        "create",
			Description: "Create resources",
			Commands: []DynamicToolCommand{
				{Name: "create_cluster", Command: "create cluster", Description: "Create a new EKS cluster"},
				{Name: "create_nodegroup", Command: "create nodegroup", Description: "Create a new nodegroup for an existing cluster"},
				{Name: "create_iamserviceaccount", Command: "create iamserviceaccount", Description: "Create an IAM service account"},
				{Name: "create_iamidentitymapping", Command: "create iamidentitymapping", Description: "Create an IAM identity mapping"},
				{Name: "create_fargateprofile", Command: "create fargateprofile", Description: "Create a Fargate profile"},
			},
		},
		{
			Name:        "get",
			Description: "Get resources",
			Commands: []DynamicToolCommand{
				{Name: "get_cluster", Command: "get cluster", Description: "Get cluster(s)"},
				{Name: "get_nodegroup", Command: "get nodegroup", Description: "Get nodegroup(s)"},
				{Name: "get_iamserviceaccount", Command: "get iamserviceaccount", Description: "Get IAM service account(s)"},
				{Name: "get_iamidentitymapping", Command: "get iamidentitymapping", Description: "Get IAM identity mapping(s)"},
				{Name: "get_fargateprofile", Command: "get fargateprofile", Description: "Get Fargate profile(s)"},
			},
		},
		{
			Name:        "update",
			Description: "Update resources",
			Commands: []DynamicToolCommand{
				{Name: "update_cluster", Command: "update cluster", Description: "Update cluster configuration"},
				{Name: "update_nodegroup", Command: "update nodegroup", Description: "Update nodegroup configuration"},
				{Name: "update_cluster-logging", Command: "update cluster-logging", Description: "Update cluster logging configuration"},
				{Name: "update_cluster-config", Command: "update cluster-config", Description: "Update cluster configuration"},
				{Name: "update_addon", Command: "update addon", Description: "Update an addon"},
				{Name: "update_accessentry", Command: "update accessentry", Description: "Update an access entry"},
			},
		},
		{
			Name:        "delete",
			Description: "Delete resources",
			Commands: []DynamicToolCommand{
				{Name: "delete_cluster", Command: "delete cluster", Description: "Delete a cluster"},
				{Name: "delete_nodegroup", Command: "delete nodegroup", Description: "Delete a nodegroup"},
				{Name: "delete_iamserviceaccount", Command: "delete iamserviceaccount", Description: "Delete an IAM service account"},
				{Name: "delete_iamidentitymapping", Command: "delete iamidentitymapping", Description: "Delete an IAM identity mapping"},
				{Name: "delete_fargateprofile", Command: "delete fargateprofile", Description: "Delete a Fargate profile"},
				{Name: "delete_addon", Command: "delete addon", Description: "Delete an addon"},
				{Name: "delete_accessentry", Command: "delete accessentry", Description: "Delete an access entry"},
				{Name: "delete_podidentityassociation", Command: "delete podidentityassociation", Description: "Delete a pod identity association"},
			},
		},
		{
			Name:        "set",
			Description: "Set resource attributes",
			Commands: []DynamicToolCommand{
				{Name: "set_labels", Command: "set labels", Description: "Set labels for a nodegroup"},
			},
		},
		{
			Name:        "unset",
			Description: "Unset resource attributes",
			Commands: []DynamicToolCommand{
				{Name: "unset_labels", Command: "unset labels", Description: "Unset labels for a nodegroup"},
			},
		},
		{
			Name:        "scale",
			Description: "Scale resources",
			Commands: []DynamicToolCommand{
				{Name: "scale_nodegroup", Command: "scale nodegroup", Description: "Scale a nodegroup"},
			},
		},
		{
			Name:        "drain",
			Description: "Drain resources",
			Commands: []DynamicToolCommand{
				{Name: "drain_nodegroup", Command: "drain nodegroup", Description: "Drain a nodegroup"},
			},
		},
		{
			Name:        "utils",
			Description: "Utility commands",
			Commands: []DynamicToolCommand{
				{Name: "utils_write-kubeconfig", Command: "utils write-kubeconfig", Description: "Write kubeconfig file"},
				{Name: "utils_describe-stacks", Command: "utils describe-stacks", Description: "Describe CloudFormation stacks"},
				{Name: "utils_update-kube-proxy", Command: "utils update-kube-proxy", Description: "Update kube-proxy addon"},
				{Name: "utils_update-aws-node", Command: "utils update-aws-node", Description: "Update aws-node addon"},
				{Name: "utils_update-coredns", Command: "utils update-coredns", Description: "Update CoreDNS addon"},
				{Name: "utils_associate-iam-oidc-provider", Command: "utils associate-iam-oidc-provider", Description: "Associate IAM OIDC provider"},
			},
		},
		{
			Name:        "upgrade",
			Description: "Upgrade resources",
			Commands: []DynamicToolCommand{
				{Name: "upgrade_cluster", Command: "upgrade cluster", Description: "Upgrade a cluster"},
				{Name: "upgrade_nodegroup", Command: "upgrade nodegroup", Description: "Upgrade a nodegroup"},
			},
		},
		{
			Name:        "enable",
			Description: "Enable features",
			Commands: []DynamicToolCommand{
				{Name: "enable_flux", Command: "enable flux", Description: "Enable Flux for a cluster"},
			},
		},
		{
			Name:        "associate",
			Description: "Associate resources",
			Commands: []DynamicToolCommand{
				{Name: "associate_identityprovider", Command: "associate identityprovider", Description: "Associate an identity provider"},
			},
		},
		{
			Name:        "disassociate",
			Description: "Disassociate resources",
			Commands: []DynamicToolCommand{
				{Name: "disassociate_identityprovider", Command: "disassociate identityprovider", Description: "Disassociate an identity provider"},
			},
		},
		{
			Name:        "register",
			Description: "Register resources",
			Commands: []DynamicToolCommand{
				{Name: "register_cluster", Command: "register cluster", Description: "Register a non-EKS cluster"},
			},
		},
		{
			Name:        "deregister",
			Description: "Deregister resources",
			Commands: []DynamicToolCommand{
				{Name: "deregister_cluster", Command: "deregister cluster", Description: "Deregister a non-EKS cluster"},
			},
		},
		{
			Name:        "info",
			Description: "Get information",
			Commands: []DynamicToolCommand{
				{Name: "info", Command: "info", Description: "Get information about eksctl"},
			},
		},
		{
			Name:        "version",
			Description: "Get version",
			Commands: []DynamicToolCommand{
				{Name: "version", Command: "version", Description: "Get version of eksctl"},
			},
		},
		{
			Name:        "completion",
			Description: "Generate shell completion",
			Commands: []DynamicToolCommand{
				{Name: "completion_bash", Command: "completion bash", Description: "Generate bash completion"},
				{Name: "completion_zsh", Command: "completion zsh", Description: "Generate zsh completion"},
				{Name: "completion_fish", Command: "completion fish", Description: "Generate fish completion"},
				{Name: "completion_powershell", Command: "completion powershell", Description: "Generate powershell completion"},
			},
		},
		{
			Name:        "help",
			Description: "Get help",
			Commands: []DynamicToolCommand{
				{Name: "help", Command: "help", Description: "Get help about eksctl commands"},
			},
		},
	}

	// Register all categories
	for _, category := range categories {
		err := RegisterDynamicToolCategory(s, category)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetCommandFromToolName extracts the eksctl command from a tool name
func GetCommandFromToolName(toolName string) string {
	// Remove "eksctl_" prefix
	name := strings.TrimPrefix(toolName, "eksctl_")
	// Replace underscores with spaces
	return strings.ReplaceAll(name, "_", " ")
}
