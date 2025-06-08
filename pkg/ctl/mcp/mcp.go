// Package mcp implements the Model Context Protocol (MCP) server functionality for eksctl
// MCP allows eksctl to be used as a tool provider for AI assistants
package mcp

import (
	"context"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"

	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
	"github.com/weaveworks/eksctl/pkg/version"
)

// Command creates the `mcp` commands
// This command starts an MCP server that provides eksctl functionality through the Model Context Protocol
func Command(_ *cmdutils.FlagGrouping) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "Start an MCP (Model Context Protocol) server",
		Long:  "Start an MCP server that provides eksctl functionality through the Model Context Protocol",
		Run: func(_ *cobra.Command, _ []string) {
			startMCPServer()
		},
		Hidden: true, // Hide this command from normal usage
	}

	return cmd
}

// startMCPServer initializes and starts the MCP server
// It creates an eksctl MCP server and connects it to stdin/stdout for communication
func startMCPServer() {
	// Create the MCP server
	mcpServer, err := newEksctlMCPServer("eksctl", version.GetVersion())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating MCP server: %v\n", err)
		os.Exit(1)
	}

	// Create a stdio server
	stdioServer := server.NewStdioServer(mcpServer)

	// Start the server
	if err := stdioServer.Listen(context.Background(), os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting MCP server: %v\n", err)
		os.Exit(1)
	}
}
