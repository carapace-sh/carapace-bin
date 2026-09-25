package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var plugins_updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Updates a plugin",
	Aliases: []string{"upgrade", "up"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugins_updateCmd).Standalone()

	plugins_updateCmd.Flags().BoolP("all", "a", false, "Update all plugins")
	plugins_updateCmd.Flags().StringP("jobs", "j", "", "Number of jobs to run in parallel")
	pluginsCmd.AddCommand(plugins_updateCmd)

	carapace.Gen(plugins_updateCmd).PositionalAnyCompletion(
		action.ActionPlugins(),
	)
}
