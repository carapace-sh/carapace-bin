package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var settings_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Append a value to an array setting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_addCmd).Standalone()

	settingsCmd.AddCommand(settings_addCmd)

	carapace.Gen(settings_addCmd).PositionalCompletion(
		action.ActionSettings(),
	)
}
