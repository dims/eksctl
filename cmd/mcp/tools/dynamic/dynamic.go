package dynamic

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
	"github.com/weaveworks/eksctl/pkg/actions/anywhere"
	"github.com/weaveworks/eksctl/pkg/version"
)

// RegisterTools registers all dynamic tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register a tool to refresh the command cache
	s.AddTool(mcp.NewTool(
		"eksctl_refresh_commands",
		mcp.WithDescription("Refresh the eksctl command cache"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		err := common.RefreshCommandCache()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to refresh command cache: %v", err)), nil
		}
		return mcp.NewToolResultText("Command cache refreshed successfully"), nil
	})

	// Register a tool to list all available commands
	s.AddTool(mcp.NewTool(
		"eksctl_list_commands",
		mcp.WithDescription("List all available eksctl commands"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		commands, err := common.ListAvailableCommands()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to list commands: %v", err)), nil
		}
		
		result := "Available eksctl commands:\n\n"
		for _, cmd := range commands {
			result += fmt.Sprintf("- %s\n", cmd)
		}
		
		return mcp.NewToolResultText(result), nil
	})

	// Register a tool to get command details
	s.AddTool(mcp.NewTool(
		"eksctl_command_info",
		mcp.WithDescription("Get detailed information about an eksctl command"),
		mcp.WithString("command",
			mcp.Description("The eksctl command to get information about"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		command := request.GetString("command", "")
		if command == "" {
			return mcp.NewToolResultError("Command is required"), nil
		}
		
		info, err := common.DiscoverCommandParameters(command)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get command info: %v", err)), nil
		}
		
		result := fmt.Sprintf("Command: %s\n", info.Command)
		result += fmt.Sprintf("Description: %s\n\n", info.Description)
		result += "Parameters:\n"
		
		for _, param := range info.Parameters {
			result += fmt.Sprintf("  --%s\n", param.Name)
			result += fmt.Sprintf("    Description: %s\n", param.Description)
			result += fmt.Sprintf("    Type: %s\n", param.Type)
			if param.Required {
				result += "    Required: Yes\n"
			}
			if len(param.Choices) > 0 {
				result += fmt.Sprintf("    Choices: %s\n", strings.Join(param.Choices, ", "))
			}
			result += "\n"
		}
		
		return mcp.NewToolResultText(result), nil
	})

	// Register the anywhere command with special handling
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

	// Register all dynamic tools
	err := common.RegisterAllDynamicTools(s)
	if err != nil {
		fmt.Printf("Error registering dynamic tools: %v\n", err)
	}
}
