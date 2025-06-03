package associate

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

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
		name := request.GetString("name", "")
		region := request.GetString("region", "")
		idpType := request.GetString("identity-provider-type", "")
		
		if name == "" || region == "" || idpType == "" {
			return mcp.NewToolResultError("name, region, and identity-provider-type are required"), nil
		}
		
		return mcp.NewToolResultText(fmt.Sprintf("Associating identity provider of type %s with cluster %s in region %s (stub implementation)", idpType, name, region)), nil
	})
}
