package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/agy"
	"github.com/spf13/cobra"
)

var pluginDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginDisableCmd).Standalone()
	pluginCmd.AddCommand(pluginDisableCmd)

	carapace.Gen(pluginDisableCmd).PositionalCompletion(
		agy.ActionPlugins(),
	)
}
