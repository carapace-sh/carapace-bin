package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/agy"
	"github.com/spf13/cobra"
)

var pluginEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginEnableCmd).Standalone()
	pluginCmd.AddCommand(pluginEnableCmd)

	carapace.Gen(pluginEnableCmd).PositionalCompletion(
		agy.ActionPlugins(),
	)
}
