package help

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RegisterTools registers the help tool with the MCP server
func RegisterTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool(
		"eksctl_help",
		mcp.WithDescription("Help about any command"),
		mcp.WithString("command",
			mcp.Description("Command to get help for"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		command := request.GetString("command", "")

		if command == "" {
			// General help
			return mcp.NewToolResultText(`eksctl - The official CLI for Amazon EKS

Usage: eksctl [command] [flags]

Commands:
  associate       Associate resources with a cluster
  completion      Generate shell completion
  create          Create resource(s)
  delete          Delete resource(s)
  deregister      Deregister a non-EKS cluster
  disassociate    Disassociate resources from a cluster
  drain           Drain resource(s)
  enable          Enable features in a cluster
  get             Get resource(s)
  help            Help about any command
  info            Output information about eksctl
  register        Register a non-EKS cluster
  scale           Scale resources(s)
  set             Set values
  unset           Unset values
  update          Update resource(s)
  upgrade         Upgrade resource(s)
  utils           Various utils
  version         Output the version of eksctl

Flags:
  -h, --help            Help for this command
      --profile string  AWS credentials profile to use (overrides the AWS_PROFILE environment variable)
  -v, --verbose int32   Set log level, use 0 to silence, 4 for debugging and 5 for debugging with AWS debug logging (default 3)

Use "eksctl [command] --help" for more information about a command.`), nil
		}

		// Command-specific help
		switch command {
		case "create":
			return mcp.NewToolResultText(`eksctl create - Create resource(s)

Usage: eksctl create [command] [flags]

Commands:
  accessentry             Create an access entry
  addon                   Create an addon
  cluster                 Create a cluster
  fargateprofile          Create a Fargate profile
  iamidentitymapping      Create an IAM identity mapping
  iamserviceaccount       Create an IAM service account
  nodegroup               Create a nodegroup
  podidentityassociation  Create a pod identity association

Flags:
  -h, --help   Help for this command

Global Flags:
      --profile string   AWS credentials profile to use (overrides the AWS_PROFILE environment variable)
  -v, --verbose int32    Set log level, use 0 to silence, 4 for debugging and 5 for debugging with AWS debug logging (default 3)

Use "eksctl create [command] --help" for more information about a command.`), nil
		case "get":
			return mcp.NewToolResultText(`eksctl get - Get resource(s)

Usage: eksctl get [command] [flags]

Commands:
  accessentry             Get access entry(s)
  addon                   Get addon(s)
  cluster                 Get cluster(s)
  fargateprofile          Get Fargate profile(s)
  iamidentitymapping      Get IAM identity mapping(s)
  iamserviceaccount       Get IAM service account(s)
  identityprovider        Get identity provider(s)
  labels                  Get labels for a nodegroup
  nodegroup               Get nodegroup(s)
  podidentityassociation  Get pod identity association(s)

Flags:
  -h, --help   Help for this command

Global Flags:
      --profile string   AWS credentials profile to use (overrides the AWS_PROFILE environment variable)
  -v, --verbose int32    Set log level, use 0 to silence, 4 for debugging and 5 for debugging with AWS debug logging (default 3)

Use "eksctl get [command] --help" for more information about a command.`), nil
		default:
			return mcp.NewToolResultText("No help found for command: " + command), nil
		}
	})
}
