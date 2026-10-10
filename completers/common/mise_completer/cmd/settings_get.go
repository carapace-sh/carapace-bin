package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var settings_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Show the effective value of a setting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_getCmd).Standalone()

	settingsCmd.AddCommand(settings_getCmd)

	carapace.Gen(settings_getCmd).PositionalCompletion(
		action.ActionSettings(),
	)
}
