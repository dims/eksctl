package create

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all create tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register create cluster command
	RegisterCreateClusterTool(s)

	// Register create nodegroup command
	RegisterCreateNodegroupTool(s)

	// Register create iamserviceaccount command
	RegisterCreateIAMServiceAccountTool(s)

	// Register create iamidentitymapping command
	RegisterCreateIAMIdentityMappingTool(s)

	// Register create fargateprofile command
	RegisterCreateFargateProfileTool(s)

	// Register create addon command
	RegisterCreateAddonTool(s)

	// Register create accessentry command
	RegisterCreateAccessEntryTool(s)

	// Register create podidentityassociation command
	RegisterCreatePodIdentityAssociationTool(s)
}
