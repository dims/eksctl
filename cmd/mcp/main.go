package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/weaveworks/eksctl/pkg/ctl/misc"
	"os"
	"os/exec"
	"strings"

	"github.com/kris-nova/logger"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/weaveworks/eksctl/pkg/ctl/associate"
	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
	"github.com/weaveworks/eksctl/pkg/ctl/completion"
	"github.com/weaveworks/eksctl/pkg/ctl/create"
	"github.com/weaveworks/eksctl/pkg/ctl/delete"
	"github.com/weaveworks/eksctl/pkg/ctl/deregister"
	"github.com/weaveworks/eksctl/pkg/ctl/disassociate"
	"github.com/weaveworks/eksctl/pkg/ctl/drain"
	"github.com/weaveworks/eksctl/pkg/ctl/enable"
	"github.com/weaveworks/eksctl/pkg/ctl/get"
	"github.com/weaveworks/eksctl/pkg/ctl/register"
	"github.com/weaveworks/eksctl/pkg/ctl/scale"
	"github.com/weaveworks/eksctl/pkg/ctl/set"
	"github.com/weaveworks/eksctl/pkg/ctl/unset"
	"github.com/weaveworks/eksctl/pkg/ctl/update"
	"github.com/weaveworks/eksctl/pkg/ctl/upgrade"
	"github.com/weaveworks/eksctl/pkg/ctl/utils"
)

func addCommands(rootCmd *cobra.Command, flagGrouping *cmdutils.FlagGrouping) {
	rootCmd.AddCommand(associate.Command(flagGrouping))
	rootCmd.AddCommand(create.Command(flagGrouping))
	rootCmd.AddCommand(disassociate.Command(flagGrouping))
	rootCmd.AddCommand(get.Command(flagGrouping))
	rootCmd.AddCommand(update.Command(flagGrouping))
	rootCmd.AddCommand(upgrade.Command(flagGrouping))
	rootCmd.AddCommand(delete.Command(flagGrouping))
	rootCmd.AddCommand(set.Command(flagGrouping))
	rootCmd.AddCommand(unset.Command(flagGrouping))
	rootCmd.AddCommand(scale.Command(flagGrouping))
	rootCmd.AddCommand(drain.Command(flagGrouping))
	rootCmd.AddCommand(enable.Command(flagGrouping))
	rootCmd.AddCommand(register.Command(flagGrouping))
	rootCmd.AddCommand(deregister.Command(flagGrouping))
	rootCmd.AddCommand(utils.Command(flagGrouping))
	rootCmd.AddCommand(completion.Command(rootCmd))
	misc.Command(flagGrouping, rootCmd)
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "eksctl [command]",
		Short: "The official CLI for Amazon EKS",
		Run: func(c *cobra.Command, _ []string) {
			if err := c.Help(); err != nil {
				logger.Debug("ignoring cobra error %q", err.Error())
			}
		},
		SilenceUsage: true,
	}
	flagGrouping := cmdutils.NewGrouping()

	addCommands(rootCmd, flagGrouping)

	// Create a new MCP server
	s := server.NewMCPServer("eksctl", "0.1.0", server.WithInstructions("MCP server for eksctl"))

	registerMCPToolRecursive(flagGrouping, rootCmd, s/*, &buf*/)

	// Create a stdio server
	stdioServer := server.NewStdioServer(s)

	// Start the server
	if err := stdioServer.Listen(context.Background(), os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting MCP server: %v\n", err)
		os.Exit(1)
	}
}

func registerMCPToolRecursive(flagGrouping *cmdutils.FlagGrouping, cmd *cobra.Command, s *server.MCPServer) {
	registerMCPTool(flagGrouping, cmd, s)
	for _, subCmd := range cmd.Commands() {
		registerMCPToolRecursive(flagGrouping, subCmd, s)
	}
}

func registerMCPTool(flagGrouping *cmdutils.FlagGrouping, cmd *cobra.Command, s *server.MCPServer) {
	// Skip commands that are not meant to be exposed as tools
	if cmd.Hidden || cmd.Name() == "help" || cmd.Name() == "completion" {
		return
	}

    var buf bytes.Buffer
	cmd.SetOut(&buf)
	err := flagGrouping.Usage(cmd)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error printing usage: %v\n", err)
		os.Exit(1)
	}

	// Build the command path (e.g., "create cluster")
	var cmdPath []string
	current := cmd
	for current != nil && current.Name() != "eksctl" {
		cmdPath = append([]string{current.Name()}, cmdPath...)
		current = current.Parent()
	}

	// Skip if no path (this is the root command)
	if len(cmdPath) == 0 {
		return
	}

	commandPath := strings.Join(cmdPath, " ")
	toolName := "eksctl_" + strings.Join(cmdPath, "_")

	// Extract description from the command
	description := cmd.Short
	if description == "" {
		description = cmd.Long
	}
	if description == "" {
		description = "Run the eksctl " + commandPath + " command"
	}

	// Create tool options starting with description
	toolOptions := []mcp.ToolOption{mcp.WithDescription(description + "\n\n" + buf.String())}

	// Add parameters based on command flags
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		// Skip deprecated flags
		if flag.Deprecated != "" {
			return
		}

		// Determine parameter type and add appropriate option
		switch flag.Value.Type() {
		case "bool":
			toolOptions = append(toolOptions, mcp.WithBoolean(
				flag.Name,
				mcp.Description(flag.Usage),
			))
		case "stringSlice", "stringArray":
			// Handle string arrays as regular strings with comma-separated values
			toolOptions = append(toolOptions, mcp.WithString(
				flag.Name,
				mcp.Description(flag.Usage + " (comma-separated values)"),
			))
		case "intSlice", "intArray":
			// Handle int arrays as regular strings with comma-separated values
			toolOptions = append(toolOptions, mcp.WithString(
				flag.Name,
				mcp.Description(flag.Usage + " (comma-separated values)"),
			))
		case "int", "int32", "int64", "float", "float32", "float64":
			// Use string for numbers as well for simplicity
			toolOptions = append(toolOptions, mcp.WithString(
				flag.Name,
				mcp.Description(flag.Usage),
			))
		default:
			// Default to string for all other types
			stringOpts := []mcp.PropertyOption{mcp.Description(flag.Usage)}

			// Mark required if the flag is required
			if flag.Annotations != nil {
				if _, required := flag.Annotations["cobra_annotation_required"]; required {
					stringOpts = append(stringOpts, mcp.Required())
				}
			}

			toolOptions = append(toolOptions, mcp.WithString(
				flag.Name,
				stringOpts...,
			))
		}
	})

	// Create the tool with all options
	tool := mcp.NewTool(toolName, toolOptions...)

	// Register the tool with the server
	s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Build arguments for the eksctl command
		args := cmdPath

		// Add flag values from the request
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			name := flag.Name

			// Handle different flag types
			switch flag.Value.Type() {
			case "bool":
				if request.GetBool(name, false) {
					args = append(args, "--"+name)
				}
			case "stringSlice", "stringArray", "intSlice", "intArray":
				// Handle arrays as comma-separated values
				if value := request.GetString(name, ""); value != "" {
					values := strings.Split(value, ",")
					for _, v := range values {
						args = append(args, "--"+name, strings.TrimSpace(v))
					}
				}
			default:
				// Handle string and number types
				if value := request.GetString(name, ""); value != "" {
					args = append(args, "--"+name, value)
				}
			}
		})

		// Execute the command
		cmd := exec.CommandContext(ctx, "eksctl", args...)

		// Capture output
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()
		if err != nil {
			errorMsg := stderr.String()
			if errorMsg == "" {
				errorMsg = err.Error()
			}
			return mcp.NewToolResultError(errorMsg), nil
		}

		return mcp.NewToolResultText(stdout.String()), nil
	})
}
