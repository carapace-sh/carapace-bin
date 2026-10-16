package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var settings_unsetCmd = &cobra.Command{
	Use:     "unset",
	Short:   "Clear a setting",
	Aliases: []string{"rm", "remove", "delete", "del"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_unsetCmd).Standalone()

	settingsCmd.AddCommand(settings_unsetCmd)

	carapace.Gen(settings_unsetCmd).PositionalCompletion(
		action.ActionSettings(),
	)
}
