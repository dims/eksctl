package update

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all update tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register update addon command
	RegisterUpdateAddonTool(s)
	
	// Register update auto-mode-config command
	RegisterUpdateAutoModeConfigTool(s)
	
	// Register update cluster command
	RegisterUpdateClusterTool(s)
	
	// Register update iamserviceaccount command
	RegisterUpdateIAMServiceAccountTool(s)
	
	// Register update nodegroup command
	RegisterUpdateNodegroupTool(s)
	
	// Register update podidentityassociation command
	RegisterUpdatePodIdentityAssociationTool(s)
}
