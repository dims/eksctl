package anywhere

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/pkg/actions/anywhere"
	"github.com/weaveworks/eksctl/pkg/version"
)

// RegisterTools registers the anywhere tool with the MCP server
func RegisterTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_anywhere",
		mcp.WithDescription("EKS Anywhere commands"),
		mcp.WithString("command",
			mcp.Description("Command to pass to eksctl-anywhere"),
			mcp.Required(),
		),
		mcp.WithString("args",
			mcp.Description("Arguments to pass to the command"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		command := request.GetString("command", "")
		args := request.GetString("args", "")

		// Check if eksctl-anywhere is available
		if _, err := exec.LookPath(anywhere.BinaryFileName); err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("%q plugin was not found on your path", anywhere.BinaryFileName)), nil
		}

		// Build command arguments
		cmdArgs := []string{command}
		if args != "" {
			cmdArgs = append(cmdArgs, strings.Split(args, " ")...)
		}

		// Create the command
		cmd := exec.CommandContext(ctx, anywhere.BinaryFileName, cmdArgs...)

		// Set environment
		cmd.Env = append(os.Environ(), fmt.Sprintf("EKSCTL_VERSION=%s", version.GetVersion()))

		// Capture output
		output, err := cmd.CombinedOutput()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Error executing command: %v\nOutput: %s", err, string(output))), nil
		}

		return mcp.NewToolResultText(string(output)), nil
	})
}
