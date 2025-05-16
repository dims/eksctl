package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// KarpenterInstallParams defines parameters for Karpenter installation
type KarpenterInstallParams struct {
	ClusterName           string `json:"clusterName"`
	Region                string `json:"region,omitempty"`
	Version               string `json:"version,omitempty"`
	CreateServiceAccount  bool   `json:"createServiceAccount,omitempty"`
	ServiceAccountRoleARN string `json:"serviceAccountRoleARN,omitempty"`
}

// KarpenterConfigureParams defines parameters for Karpenter configuration
type KarpenterConfigureParams struct {
	ClusterName            string `json:"clusterName"`
	Region                 string `json:"region,omitempty"`
	DefaultInstanceProfile string `json:"defaultInstanceProfile,omitempty"`
	InterruptionHandling   bool   `json:"interruptionHandling,omitempty"`
	IsolatedVPC            bool   `json:"isolatedVPC,omitempty"`
}

func registerKarpenterTools(s *server.MCPServer) {
	// Register karpenter_install tool
	s.AddTool(
		mcp.NewTool("karpenter_install", mcp.WithDescription("Install Karpenter for cluster autoscaling")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p KarpenterInstallParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			// eksctl doesn't have a direct command for Karpenter installation
			// We'll use a combination of eksctl and helm commands

			// First, create the required IAM resources
			args := []string{"create", "iamserviceaccount",
				"--cluster", p.ClusterName,
				"--name", "karpenter",
				"--namespace", "karpenter",
				"--attach-policy-arn", "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
				"--approve"}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}

			if p.ServiceAccountRoleARN != "" {
				args = append(args, "--role-arn", p.ServiceAccountRoleARN)
			} else if p.CreateServiceAccount {
				// Create a role with necessary permissions
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to create IAM service account for Karpenter: %s\n%s", err, string(output))), nil
			}

			// Now install Karpenter using Helm
			helmArgs := []string{"install", "karpenter", "karpenter/karpenter",
				"--namespace", "karpenter",
				"--create-namespace",
				"--set", fmt.Sprintf("serviceAccount.annotations.eks\\.amazonaws\\.com/role-arn=%s", p.ServiceAccountRoleARN),
				"--set", fmt.Sprintf("clusterName=%s", p.ClusterName),
				"--set", fmt.Sprintf("clusterEndpoint=$(aws eks describe-cluster --name %s --query \"cluster.endpoint\" --output text)", p.ClusterName)}

			if p.Version != "" {
				helmArgs = append(helmArgs, "--version", p.Version)
			}

			if p.Region != "" {
				helmArgs = append(helmArgs, "--set", fmt.Sprintf("aws.defaultInstanceProfile=KarpenterNodeInstanceProfile-%s", p.ClusterName))
			}

			helmCmd := exec.Command("helm", helmArgs...)
			helmOutput, err := helmCmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to install Karpenter: %s\n%s", err, string(helmOutput))), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Karpenter installation completed successfully.\n\nService Account Creation:\n%s\n\nHelm Installation:\n%s", string(output), string(helmOutput))), nil
		},
	)

	// Register karpenter_configure tool
	s.AddTool(
		mcp.NewTool("karpenter_configure", mcp.WithDescription("Configure Karpenter settings")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p KarpenterConfigureParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			// Create a ConfigMap with Karpenter configuration
			configMapYAML := fmt.Sprintf(`apiVersion: v1
kind: ConfigMap
metadata:
  name: karpenter-config
  namespace: karpenter
data:
  aws:
    clusterName: %s
`, p.ClusterName)

			if p.DefaultInstanceProfile != "" {
				configMapYAML += fmt.Sprintf("    defaultInstanceProfile: %s\n", p.DefaultInstanceProfile)
			}

			if p.InterruptionHandling {
				configMapYAML += "    interruptionHandling: true\n"
			}

			if p.IsolatedVPC {
				configMapYAML += "    isolatedVPC: true\n"
			}

			// Write the ConfigMap to a temporary file
			tempFile := "/tmp/karpenter-config.yaml"
			cmd := exec.Command("bash", "-c", fmt.Sprintf("cat > %s << 'EOF'\n%s\nEOF", tempFile, configMapYAML))
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to create config file: %s\n%s", err, string(output))), nil
			}

			// Apply the ConfigMap
			kubectlArgs := []string{"apply", "-f", tempFile}
			kubectlCmd := exec.Command("kubectl", kubectlArgs...)
			kubectlOutput, err := kubectlCmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to apply Karpenter configuration: %s\n%s", err, string(kubectlOutput))), nil
			}

			// Restart Karpenter to apply the new configuration
			restartArgs := []string{"rollout", "restart", "deployment", "karpenter", "-n", "karpenter"}
			restartCmd := exec.Command("kubectl", restartArgs...)
			restartOutput, err := restartCmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to restart Karpenter: %s\n%s", err, string(restartOutput))), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Karpenter configuration updated successfully.\n\nConfiguration:\n%s\n\nApply Output:\n%s\n\nRestart Output:\n%s", configMapYAML, string(kubectlOutput), string(restartOutput))), nil
		},
	)
}
