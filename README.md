# eksctl MCP Server

This is an implementation of the Model Context Protocol (MCP) server for eksctl. It allows eksctl commands to be executed through the MCP protocol.

## Building

```bash
go build -o eksctl-mcp ./cmd/mcp
```

## Running

```bash
./eksctl-mcp
```

## Available Tools

The following eksctl commands are available as MCP tools:

### Create Commands
- eksctl_create_cluster
- eksctl_create_nodegroup
- eksctl_create_iamserviceaccount
- eksctl_create_iamidentitymapping
- eksctl_create_fargateprofile
- eksctl_create_addon
- eksctl_create_accessentry
- eksctl_create_podidentityassociation

### Delete Commands
- eksctl_delete_cluster
- eksctl_delete_nodegroup
- eksctl_delete_iamserviceaccount
- eksctl_delete_iamidentitymapping
- eksctl_delete_fargateprofile
- eksctl_delete_addon
- eksctl_delete_accessentry
- eksctl_delete_podidentityassociation

### Get Commands
- eksctl_get_cluster
- eksctl_get_nodegroup
- eksctl_get_iamserviceaccount
- eksctl_get_iamidentitymapping
- eksctl_get_fargateprofile
- eksctl_get_addon
- eksctl_get_accessentry
- eksctl_get_podidentityassociation

### Utils Commands
- eksctl_utils_write-kubeconfig
- eksctl_utils_describe-stacks
- eksctl_utils_update-kube-proxy
- eksctl_utils_update-aws-node
- eksctl_utils_update-coredns
- eksctl_utils_associate-iam-oidc-provider

### Other Commands
- eksctl_version

## Usage with Amazon Q CLI

Once the MCP server is running, you can use it with the Amazon Q CLI:

```bash
q chat --mcp-server localhost:8080
```

Then you can use eksctl commands through the Amazon Q CLI:

```
> Create an EKS cluster in us-west-2 with 2 nodes
```

The Amazon Q CLI will use the MCP server to execute the eksctl commands.
