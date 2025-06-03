package disassociate

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all disassociate tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register disassociate identityprovider command
	RegisterDisassociateIdentityProviderTool(s)
}

// RegisterDisassociateIdentityProviderTool registers the eksctl_disassociate_identityprovider tool with the MCP server
func RegisterDisassociateIdentityProviderTool(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_disassociate_identityprovider",
		mcp.WithDescription("Disassociate an identity provider from an EKS cluster"),
		mcp.WithString("name",
			mcp.Description("EKS cluster name"),
			mcp.Required(),
		),
		mcp.WithString("region",
			mcp.Description("AWS region"),
			mcp.Required(),
		),
		mcp.WithString("identity-provider-name",
			mcp.Description("Name of identity provider"),
			mcp.Required(),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		idpName := request.GetString("identity-provider-name", "")

		if name == "" || region == "" || idpName == "" {
			return mcp.NewToolResultError("name, region, and identity-provider-name are required"), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Disassociating identity provider %s from cluster %s in region %s (stub implementation)", idpName, name, region)), nil
	})
}
