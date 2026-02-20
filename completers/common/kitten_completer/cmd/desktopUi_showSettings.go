package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showSettingsCmd = &cobra.Command{
	Use:   "show-settings",
	Short: "Print the current values of the desktop settings",
}

func init() {
	showSettingsCmd.AddCommand(desktopUiCmd)
	carapace.Gen(showSettingsCmd).Standalone()

	showSettingsCmd.Flags().Bool("allow-other-backends", false, "Normally, after printing the settings, if the settings did not come from the desktop-ui kitten the command prints an error and exits. This prevents that.")
	showSettingsCmd.Flags().Bool("as-json", false, "Show the settings as JSON for machine consumption")
	showSettingsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	showSettingsCmd.Flags().String("in-namespace", "", "Show only settings in the specified namespaces")

	carapace.Gen(showSettingsCmd).FlagCompletion(carapace.ActionMap{})
}
