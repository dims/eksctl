package tools

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
)

// ExecuteEksctlCommand executes the eksctl binary with the given arguments
func ExecuteEksctlCommand(ctx context.Context, args []string) (*mcp.CallToolResult, error) {
	// Create a context with a 45-second timeout for output collection
	timeoutCtx, cancel := context.WithTimeout(ctx, 45*time.Second)
	defer cancel()

	// Create the command with the parent context to avoid killing the process
	cmd := exec.Command("eksctl", args...)

	// Set up pipes to capture output
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return mcp.NewToolResultError("Failed to create stdout pipe: " + err.Error()), nil
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return mcp.NewToolResultError("Failed to create stderr pipe: " + err.Error()), nil
	}

	// Buffers to store output
	var stdoutBuf, stderrBuf bytes.Buffer
	
	// Create a WaitGroup to wait for the goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)
	
	// Start goroutines to read from pipes
	go func() {
		defer wg.Done()
		_, _ = stdoutBuf.ReadFrom(stdoutPipe)
	}()
	
	go func() {
		defer wg.Done()
		_, _ = stderrBuf.ReadFrom(stderrPipe)
	}()

	// Start the command in a goroutine
	errChan := make(chan error, 1)
	go func() {
		// Start the command
		if err := cmd.Start(); err != nil {
			errChan <- err
			return
		}
		
		// Wait for the command to complete
		errChan <- cmd.Wait()
	}()

	// Wait for either the command to finish or the context to timeout
	var cmdErr error
	select {
	case cmdErr = <-errChan:
		// Command completed
		wg.Wait() // Wait for output goroutines to finish
		
		// If there was an error, return the stderr output
		if cmdErr != nil {
			errMsg := stderrBuf.String()
			if errMsg == "" {
				errMsg = cmdErr.Error()
			}
			return mcp.NewToolResultError(errMsg), nil
		}
		
		// Return the stdout output
		return mcp.NewToolResultText(stdoutBuf.String()), nil
		
	case <-timeoutCtx.Done():
		if timeoutCtx.Err() == context.DeadlineExceeded {
			// Command is still running but we've reached our response time limit
			// Return the stdout collected so far without killing the process
			return mcp.NewToolResultText(stdoutBuf.String() + "\n\n[Command is still running in the background]"), nil
		}
		cmdErr = timeoutCtx.Err()
		wg.Wait()
		return mcp.NewToolResultError(cmdErr.Error()), nil
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
