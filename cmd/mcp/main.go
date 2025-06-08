package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/pkg/mcp"
)

func main() {
	// Create the MCP server
	mcpServer, err := mcp.NewEksctlMCPServer("eksctl", "0.1.0")
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
