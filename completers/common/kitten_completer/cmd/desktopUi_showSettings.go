package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var desktopUi_showSettingsCmd = &cobra.Command{
	Use:   "show-settings",
	Short: "Print the current values of the desktop settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(desktopUi_showSettingsCmd).Standalone()

	desktopUi_showSettingsCmd.Flags().Bool("allow-other-backends", false, "Normally, after printing the settings, if the settings did not come from the desktop-ui kitten the command prints an error and exits. This prevents that.")
	desktopUi_showSettingsCmd.Flags().Bool("as-json", false, "Show the settings as JSON for machine consumption")
	desktopUi_showSettingsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	desktopUi_showSettingsCmd.Flags().String("in-namespace", "", "Show only settings in the specified namespaces")
	desktopUiCmd.AddCommand(desktopUi_showSettingsCmd)

}
