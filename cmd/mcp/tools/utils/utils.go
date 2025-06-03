package utils

import (
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers all utils tools with the MCP server
func RegisterTools(s *server.MCPServer) {
	// Register utils write-kubeconfig command
	RegisterWriteKubeconfigTool(s)

	// Register utils describe-stacks command
	RegisterDescribeStacksTool(s)

	// Register utils update-kube-proxy command
	RegisterUpdateKubeProxyTool(s)

	// Register utils update-aws-node command
	RegisterUpdateAWSNodeTool(s)

	// Register utils update-coredns command
	RegisterUpdateCoreDNSTool(s)

	// Register utils associate-iam-oidc-provider command
	RegisterAssociateIAMOIDCProviderTool(s)
	
	// Register utils describe-addon-configuration command
	RegisterDescribeAddonConfigurationTool(s)
	
	// Register utils describe-addon-versions command
	RegisterDescribeAddonVersionsTool(s)
	
	// Register utils describe-cluster-versions command
	RegisterDescribeClusterVersionsTool(s)
	
	// Register utils enable-secrets-encryption command
	RegisterEnableSecretsEncryptionTool(s)
	
	// Register utils install-vpc-controllers command
	RegisterInstallVPCControllersTool(s)
	
	// Register utils migrate-to-access-entry command
	RegisterMigrateToAccessEntryTool(s)
	
	// Register utils migrate-to-pod-identity command
	RegisterMigrateToPodIdentityTool(s)
	
	// Register utils nodegroup-health command
	RegisterNodegroupHealthTool(s)
	
	// Register utils schema command
	RegisterSchemaTool(s)
	
	// Register utils update-authentication-mode command
	RegisterUpdateAuthenticationModeTool(s)
	
	// Register utils update-cluster-logging command
	RegisterUpdateClusterLoggingTool(s)
	
	// Register utils update-cluster-vpc-config command
	RegisterUpdateClusterVPCConfigTool(s)
	
	// Register utils update-legacy-subnet-settings command
	RegisterUpdateLegacySubnetSettingsTool(s)
	
	// Register utils update-zonal-shift-config command
	RegisterUpdateZonalShiftConfigTool(s)
}
