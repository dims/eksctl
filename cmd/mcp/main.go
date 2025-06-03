package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer("eksctl", "0.1.0", server.WithInstructions("MCP server for eksctl"))

	// Register all tools
	tools.RegisterTools(s)

	// Create a stdio server
	stdioServer := server.NewStdioServer(s)

	// Start the server
	if err := stdioServer.Listen(context.Background(), os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting MCP server: %v\n", err)
		os.Exit(1)
	}
}
