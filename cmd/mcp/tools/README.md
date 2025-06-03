# eksctl MCP Tools

This directory contains the implementation of the Model Context Protocol (MCP) tools for eksctl.

## Overview

The eksctl MCP tools provide a way to interact with eksctl through the Model Context Protocol, allowing AI assistants to execute eksctl commands on behalf of users.

## Architecture

The implementation uses a dynamic parameter discovery approach:

1. **Parameter Auto-Discovery**: The system automatically discovers command parameters by analyzing:
   - Output of `eksctl --help` for each command
   - Output of `eksctl completion bash` for additional parameter information

2. **Caching**: Discovered parameters are cached to avoid repeatedly running the commands.

3. **Dynamic Tool Registration**: Tools are registered with the MCP server using the discovered parameters.

## Files

- `discovery.go`: Implements parameter extraction and caching
- `dynamic.go`: Defines tool categories and registration helpers
- `eksctl.go`: Handles execution of eksctl commands
- `tools.go`: Main entry point for registering tools

## Special Cases

- **EKS Anywhere**: The `eksctl_anywhere` command has special handling because it executes a separate binary (`eksctl-anywhere`).

## Adding New Commands

When new commands are added to eksctl, they will be automatically discovered and exposed through the MCP interface. No code changes are required unless the command needs special handling.

## Maintenance

To refresh the command cache, use the `eksctl_refresh_commands` tool.

To see all available commands, use the `eksctl_list_commands` tool.

To get detailed information about a command, use the `eksctl_command_info` tool with the command name.
