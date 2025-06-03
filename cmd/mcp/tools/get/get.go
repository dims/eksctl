package get

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all get tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register get cluster command
	RegisterGetClusterTool(s)

	// Register get nodegroup command
	RegisterGetNodegroupTool(s)

	// Register get iamserviceaccount command
	RegisterGetIAMServiceAccountTool(s)

	// Register get iamidentitymapping command
	RegisterGetIAMIdentityMappingTool(s)

	// Register get fargateprofile command
	RegisterGetFargateProfileTool(s)

	// Register get addon command
	RegisterGetAddonTool(s)

	// Register get accessentry command
	RegisterGetAccessEntryTool(s)

	// Register get podidentityassociation command
	RegisterGetPodIdentityAssociationTool(s)
}
