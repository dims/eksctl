package delete

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all delete tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register delete cluster command
	RegisterDeleteClusterTool(s)

	// Register delete nodegroup command
	RegisterDeleteNodegroupTool(s)

	// Register delete iamserviceaccount command
	RegisterDeleteIAMServiceAccountTool(s)

	// Register delete iamidentitymapping command
	RegisterDeleteIAMIdentityMappingTool(s)

	// Register delete fargateprofile command
	RegisterDeleteFargateProfileTool(s)

	// Register delete addon command
	RegisterDeleteAddonTool(s)

	// Register delete accessentry command
	RegisterDeleteAccessEntryTool(s)

	// Register delete podidentityassociation command
	RegisterDeletePodIdentityAssociationTool(s)
}
