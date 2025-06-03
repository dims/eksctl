package tools

import (
	"bytes"
	"context"
	"os/exec"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

// ExecuteEksctlCommand executes the eksctl binary with the given arguments
func ExecuteEksctlCommand(ctx context.Context, args []string) (*mcp.CallToolResult, error) {
	// Create the command
	cmd := exec.CommandContext(ctx, "eksctl", args...)

	// Capture stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute the command
	err := cmd.Run()

	// If there was an error, return the stderr output
	if err != nil {
		errMsg := stderr.String()
		if errMsg == "" {
			errMsg = err.Error()
		}
		return mcp.NewToolResultError(errMsg), nil
	}

	// Return the stdout output
	return mcp.NewToolResultText(stdout.String()), nil
}

// BuildEksctlArgs builds the arguments for the eksctl command from the request parameters
func BuildEksctlArgs(command string, request mcp.CallToolRequest) []string {
	// Start with the command
	args := strings.Split(command, " ")

	// Process common parameters that might be present
	// This is a simplified approach since we can't iterate through all parameters
	commonParams := []string{
		"name", "region", "cluster", "nodegroup", "addon", "namespace", "service-account",
		"profile", "output", "version", "timeout", "approve", "wait", "force",
		"arn", "username", "principal-arn", "role-arn", "config-file",
		"max-pods-per-node", "kubernetes-groups", "type", "access-policy-arn", "access-scope",
		"nodes", "nodes-min", "nodes-max", "node-type", "node-volume-size", "node-volume-type",
		"ssh-access", "ssh-public-key", "vpc-cidr", "without-nodegroup", "dry-run", "tags",
		"zones", "with-oidc", "fargate", "profile-name", "cfn-role-arn", "cfn-disable-rollback",
		"managed", "spot", "node-private-networking", "asg-access", "external-dns-access",
		"full-ecr-access", "alb-ingress-access", "attach-policy-arn", "attach-policy",
		"role-name", "role-only", "labels", "selectors-file", "service-account-role-arn",
		"configuration-values", "resolve-conflicts", "preserve", "disable-eviction", "parallel",
		"pod-eviction-wait-period", "enable-types", "disable-types", "all",
	}

	// Add parameters that have values
	for _, param := range commonParams {
		value := request.GetString(param, "")
		if value != "" {
			args = append(args, "--"+param, value)
		}
	}

	// Handle boolean flags separately
	boolFlags := []string{
		"approve", "wait", "force", "dry-run", "with-oidc", "fargate", "cfn-disable-rollback",
		"managed", "spot", "node-private-networking", "asg-access", "external-dns-access",
		"full-ecr-access", "alb-ingress-access", "role-only", "preserve", "disable-eviction",
		"without-nodegroup", "all",
	}

	for _, flag := range boolFlags {
		if request.GetBool(flag, false) {
			args = append(args, "--"+flag)
		}
	}

	return args
}

// ExecuteEksctlCommandFromRequest builds arguments from the request and executes the eksctl command
func ExecuteEksctlCommandFromRequest(ctx context.Context, command string, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := BuildEksctlArgs(command, request)
	return ExecuteEksctlCommand(ctx, args)
}
