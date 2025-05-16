# MCP Tools for eksctl

This directory contains the implementation of Model Context Protocol (MCP) tools for eksctl. These tools allow Amazon Q to interact with eksctl functionality through natural language.

## Implementation Status

This is a proof of concept implementation that demonstrates how to integrate eksctl with Amazon Q using the MCP protocol. We've successfully added the `github.com/mark3labs/mcp-go` package, but there are still API compatibility issues to resolve.

## Current Status

We've successfully:
1. Added the MCP dependency to the project
2. Created the server initialization code using stdio for communication
3. Implemented the cluster management tools using the latest MCP API

However, we still need to update the remaining tool implementations to match the current MCP API. The main changes needed are:

1. Update all tool registrations to use `mcp.NewTool()` with functional options
2. Replace the old parameter definitions with JSON schema
3. Update the handler functions to use the new `mcp.Context` parameter and return `*mcp.CallToolResult`

## Tool Categories

The implementation includes the following tool categories:

1. **Cluster Management Tools** (Implemented with latest API)
   - `cluster_create`: Create new EKS clusters
   - `cluster_delete`: Delete existing clusters
   - `cluster_get`: Get information about clusters
   - `cluster_update`: Update cluster configurations
   - `cluster_describe`: Get detailed information about clusters

2. **Node Group Management Tools** (Needs API update)
   - `nodegroup_create`: Create node groups
   - `nodegroup_delete`: Delete node groups
   - `nodegroup_scale`: Scale node groups
   - `nodegroup_update`: Update node group configurations
   - `nodegroup_list`: List all node groups in a cluster

3. **Add-on Management Tools** (Needs API update)
   - `addon_install`: Install EKS add-ons
   - `addon_update`: Update EKS add-ons
   - `addon_delete`: Remove EKS add-ons
   - `addon_list`: List available and installed add-ons

4. **Karpenter Integration Tools** (Needs API update)
   - `karpenter_install`: Install Karpenter for cluster autoscaling
   - `karpenter_configure`: Configure Karpenter settings

5. **Auto Mode Management Tools** (Needs API update)
   - `automode_enable`: Enable EKS Auto Mode
   - `automode_configure`: Configure Auto Mode settings

6. **Utility Tools** (Needs API update)
   - `kubeconfig_get`: Generate or update kubeconfig for cluster access
   - `utils_describe_stacks`: Describe CloudFormation stacks for troubleshooting
   - `utils_update_cluster_logging`: Update cluster logging configurations

## Usage

Once all the tools are updated to use the latest MCP API, users will be able to start the MCP server by running:

```bash
eksctl mcp
```

This will start the MCP server using stdio (standard input/output) for communication with Amazon Q.

## Next Steps

To complete this implementation:

1. Update the remaining tool implementations to use the latest MCP API
2. Test each tool to ensure it works correctly
3. Add error handling and logging
4. Test the integration with Amazon Q
