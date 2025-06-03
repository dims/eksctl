package associate

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/weaveworks/eksctl/cmd/mcp/tools/common"
)

// RegisterTools registers all associate tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register associate identityprovider command
	RegisterAssociateIdentityProviderTool(s)
}

// RegisterAssociateIdentityProviderTool registers the eksctl_associate_identityprovider tool with the MCP server
func RegisterAssociateIdentityProviderTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_associate_identityprovider",
		mcp.WithDescription("Associate an identity provider with an EKS cluster"),
		mcp.WithString("name",
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("identity-provider-type",
			mcp.Description("Type of identity provider"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return common.ExecuteEksctlCommandFromRequest(ctx, "associate identityprovider", request)
	})
}
