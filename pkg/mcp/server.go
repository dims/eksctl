package mcp

import (
	"context"
	"os"

	"github.com/mark3labs/mcp-go/server"
)

// StartMCPServer starts the MCP server with all eksctl tools
func StartMCPServer() error {
	// Create a new MCP server with stdio transport
	s := server.NewMCPServer("eksctl", "1.0.0")
	
	// Set up stdio transport
	stdioServer := server.NewStdioServer(s)

	// Register all tool categories
	registerClusterTools(s)
	registerNodeGroupTools(s)
	registerAddonTools(s)
	registerKarpenterTools(s)
	registerAutoModeTools(s)
	registerUtilityTools(s)

	// Start the server with stdio transport
	ctx := context.Background()
	return stdioServer.Listen(ctx, os.Stdin, os.Stdout)
}
