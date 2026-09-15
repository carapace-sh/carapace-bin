package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var settings_setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Add/update a setting",
	Aliases: []string{"create"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_setCmd).Standalone()

	settingsCmd.AddCommand(settings_setCmd)

	carapace.Gen(settings_setCmd).PositionalCompletion(
		action.ActionSettings(),
	)
}
