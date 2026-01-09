package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var settings_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets the value of an admin setting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_setCmd).Standalone()

	settings_setCmd.Flags().String("setting", "", "Name of the setting to modify")
	settings_setCmd.Flags().String("value", "", "Value to set for the setting.")
	settingsCmd.AddCommand(settings_setCmd)

	// TODO settings: https://github.com/microsoft/winget-cli/blob/master/doc/Settings.md
}
