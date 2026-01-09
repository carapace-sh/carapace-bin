package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var settings_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets an admin setting to its default value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_resetCmd).Standalone()

	settings_resetCmd.Flags().String("setting", "", "Name of the setting to modify")
	settingsCmd.AddCommand(settings_resetCmd)
}
