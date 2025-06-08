package misc

import (
	"github.com/spf13/cobra"
	"github.com/weaveworks/eksctl/pkg/ctl/cmdutils"
)

func Command(flagGrouping *cmdutils.FlagGrouping, rootCmd *cobra.Command) *cobra.Command {
	cmdutils.AddResourceCmd(flagGrouping, rootCmd, infoCmd)
	cmdutils.AddResourceCmd(flagGrouping, rootCmd, versionCmd)
	return rootCmd
}

