package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Manage settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settingsCmd).Standalone()

	settingsCmd.Flags().BoolP("all", "a", false, "List all settings")
	settingsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	settingsCmd.Flags().BoolP("local", "l", false, "Use local config file")
	settingsCmd.Flags().BoolP("toml", "T", false, "Output in TOML format")
	rootCmd.AddCommand(settingsCmd)

	carapace.Gen(settingsCmd).PositionalCompletion(
		action.ActionSettings(),
	)
}
