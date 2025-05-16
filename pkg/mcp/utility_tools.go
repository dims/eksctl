package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// KubeconfigGetParams defines parameters for getting kubeconfig
type KubeconfigGetParams struct {
	ClusterName string `json:"clusterName"`
	Region      string `json:"region,omitempty"`
	RoleARN     string `json:"roleARN,omitempty"`
	Profile     string `json:"profile,omitempty"`
	SetContext  bool   `json:"setContext,omitempty"`
}

// UtilsDescribeStacksParams defines parameters for describing CloudFormation stacks
type UtilsDescribeStacksParams struct {
	ClusterName string `json:"clusterName"`
	Region      string `json:"region,omitempty"`
	StackName   string `json:"stackName,omitempty"`
	Output      string `json:"output,omitempty"`
}

// UtilsUpdateClusterLoggingParams defines parameters for updating cluster logging
type UtilsUpdateClusterLoggingParams struct {
	ClusterName  string   `json:"clusterName"`
	Region       string   `json:"region,omitempty"`
	EnableTypes  []string `json:"enableTypes,omitempty"`
	DisableTypes []string `json:"disableTypes,omitempty"`
}

func registerUtilityTools(s *server.MCPServer) {
	// Register kubeconfig_get tool
	s.AddTool(
		mcp.NewTool("kubeconfig_get", mcp.WithDescription("Generate or update kubeconfig for cluster access")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p KubeconfigGetParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"utils", "write-kubeconfig", "--cluster", p.ClusterName}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.RoleARN != "" {
				args = append(args, "--role-arn", p.RoleARN)
			}
			if p.Profile != "" {
				args = append(args, "--profile", p.Profile)
			}
			if p.SetContext {
				args = append(args, "--set-context")
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to get kubeconfig: %s\n%s", err, string(output))), nil
			}

			// Also get the kubeconfig content to return
			kubeconfigCmd := exec.Command("kubectl", "config", "view", "--minify", "--flatten")
			kubeconfigOutput, err := kubeconfigCmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultText(string(output)), nil // Return just the eksctl output if kubectl fails
			}

			return mcp.NewToolResultText(fmt.Sprintf("%s\n\nKubeconfig:\n%s", string(output), string(kubeconfigOutput))), nil
		},
	)

	// Register utils_describe_stacks tool
	s.AddTool(
		mcp.NewTool("utils_describe_stacks", mcp.WithDescription("Describe CloudFormation stacks for troubleshooting")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p UtilsDescribeStacksParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"utils", "describe-stacks", "--cluster", p.ClusterName}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if p.StackName != "" {
				args = append(args, "--stack-name", p.StackName)
			}
			if p.Output != "" {
				args = append(args, "--output", p.Output)
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to describe stacks: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)

	// Register utils_update_cluster_logging tool
	s.AddTool(
		mcp.NewTool("utils_update_cluster_logging", mcp.WithDescription("Update cluster logging configurations")),
		func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			var p UtilsUpdateClusterLoggingParams
			if err := json.Unmarshal([]byte(fmt.Sprintf("%v", request.Params.Arguments)), &p); err != nil {
				return mcp.NewToolResultError("Failed to parse parameters"), err
			}

			args := []string{"utils", "update-cluster-logging", "--cluster", p.ClusterName}

			if p.Region != "" {
				args = append(args, "--region", p.Region)
			}
			if len(p.EnableTypes) > 0 {
				args = append(args, "--enable-types", strings.Join(p.EnableTypes, ","))
			}
			if len(p.DisableTypes) > 0 {
				args = append(args, "--disable-types", strings.Join(p.DisableTypes, ","))
			}

			cmd := exec.Command("eksctl", args...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to update cluster logging: %s\n%s", err, string(output))), nil
			}

			return mcp.NewToolResultText(string(output)), nil
		},
	)
}
