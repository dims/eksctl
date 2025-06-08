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
func Command(_ *cmdutils.FlagGrouping) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mcp",
		Short: "Start an MCP (Model Context Protocol) server",
		Long:  "Start an MCP server that provides eksctl functionality through the Model Context Protocol",
		Run: func(_ *cobra.Command, _ []string) {
			startMCPServer()
		},
		Hidden: true,
	}

	return cmd
}

func startMCPServer() {
	// Create the MCP server
	mcpServer, err := NewEksctlMCPServer("eksctl", version.GetVersion())
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
