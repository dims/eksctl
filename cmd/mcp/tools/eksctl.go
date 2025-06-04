package tools

import (
	"bytes"
	"context"
	"io"
	"os/exec"
	"strings"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

func ExecuteEksctlCommand(ctx context.Context, args []string) (*mcp.CallToolResult, error) {
	// Create a context for output collection with a 45-second timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, 45*time.Second)
	defer cancel()

	cmd := exec.Command("eksctl", args...)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return mcp.NewToolResultError("Failed to create stdout pipe: " + err.Error()), nil
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return mcp.NewToolResultError("Failed to create stderr pipe: " + err.Error()), nil
	}

	if err := cmd.Start(); err != nil {
		return mcp.NewToolResultError("Failed to start eksctl: " + err.Error()), nil
	}

	// Channels to collect output
	stdoutCh := make(chan string, 1)
	stderrCh := make(chan string, 1)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, stdoutPipe)
		stdoutCh <- buf.String()
	}()
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, stderrPipe)
		stderrCh <- buf.String()
	}()

	// Wait for either timeout or command completion
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-timeoutCtx.Done():
		// Timeout: collect whatever output is available
		stdout := ""
		select {
		case stdout = <-stdoutCh:
		default:
		}
		return mcp.NewToolResultText(stdout + "\n\n[Command is still running in the background]"), nil
	case err := <-done:
		// Command finished: collect output
		stdout := <-stdoutCh
		stderr := <-stderrCh
		if err != nil {
			if stderr == "" {
				stderr = err.Error()
			}
			return mcp.NewToolResultError(stderr), nil
		}
		return mcp.NewToolResultText(stdout), nil
	}
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
