package info

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/pkg/version"
)

// RegisterTools registers the info tool with the MCP server
func RegisterTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_info",
		mcp.WithDescription("Display information about eksctl"),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		versionInfo := version.GetVersionInfo()
		info := fmt.Sprintf(`eksctl version: %s
Go version: go1.21.0
Git commit: %s
Build date: %s
`, version.GetVersion(), versionInfo.Metadata.GitCommit, versionInfo.Metadata.BuildDate)
		return mcp.NewToolResultText(info), nil
	})
}
