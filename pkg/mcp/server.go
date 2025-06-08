package mcp

import (
	"github.com/kris-nova/logger"
	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"

	"github.com/weaveworks/eksctl/pkg/ctl/associate"
	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
	"github.com/weaveworks/eksctl/pkg/ctl/completion"
	"github.com/weaveworks/eksctl/pkg/ctl/create"
	"github.com/weaveworks/eksctl/pkg/ctl/delete"
	"github.com/weaveworks/eksctl/pkg/ctl/deregister"
	"github.com/weaveworks/eksctl/pkg/ctl/disassociate"
	"github.com/weaveworks/eksctl/pkg/ctl/drain"
	"github.com/weaveworks/eksctl/pkg/ctl/enable"
	"github.com/weaveworks/eksctl/pkg/ctl/get"
	"github.com/weaveworks/eksctl/pkg/ctl/misc"
	"github.com/weaveworks/eksctl/pkg/ctl/register"
	"github.com/weaveworks/eksctl/pkg/ctl/scale"
	"github.com/weaveworks/eksctl/pkg/ctl/set"
	"github.com/weaveworks/eksctl/pkg/ctl/unset"
	"github.com/weaveworks/eksctl/pkg/ctl/update"
	"github.com/weaveworks/eksctl/pkg/ctl/upgrade"
	"github.com/weaveworks/eksctl/pkg/ctl/utils"
)

// NewEksctlMCPServer creates and configures an MCP server for eksctl
func NewEksctlMCPServer(name, version string) (*server.MCPServer, error) {
	rootCmd := createRootCommand()
	flagGrouping := cmdutils.NewGrouping()

	addCommands(rootCmd, flagGrouping)

	// Create a new MCP server
	s := server.NewMCPServer(name, version, server.WithInstructions("MCP server for eksctl"))

	// Register all tools
	if err := registerTools(s, rootCmd, flagGrouping); err != nil {
		return nil, err
	}

	return s, nil
}

// createRootCommand creates the root eksctl command
func createRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "eksctl-mcp [command]",
		Short: "An MCP Server for Amazon EKS",
		Run: func(c *cobra.Command, _ []string) {
			if err := c.Help(); err != nil {
				logger.Debug("ignoring cobra error %q", err.Error())
			}
		},
		SilenceUsage: true,
	}
	return rootCmd
}

// addCommands adds all eksctl subcommands to the root command
func addCommands(rootCmd *cobra.Command, flagGrouping *cmdutils.FlagGrouping) {
	rootCmd.AddCommand(associate.Command(flagGrouping))
	rootCmd.AddCommand(create.Command(flagGrouping))
	rootCmd.AddCommand(disassociate.Command(flagGrouping))
	rootCmd.AddCommand(get.Command(flagGrouping))
	rootCmd.AddCommand(update.Command(flagGrouping))
	rootCmd.AddCommand(upgrade.Command(flagGrouping))
	rootCmd.AddCommand(delete.Command(flagGrouping))
	rootCmd.AddCommand(set.Command(flagGrouping))
	rootCmd.AddCommand(unset.Command(flagGrouping))
	rootCmd.AddCommand(scale.Command(flagGrouping))
	rootCmd.AddCommand(drain.Command(flagGrouping))
	rootCmd.AddCommand(enable.Command(flagGrouping))
	rootCmd.AddCommand(register.Command(flagGrouping))
	rootCmd.AddCommand(deregister.Command(flagGrouping))
	rootCmd.AddCommand(utils.Command(flagGrouping))
	rootCmd.AddCommand(completion.Command(rootCmd))
	misc.Command(flagGrouping, rootCmd)
}
