package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var privacyCmd = &cobra.Command{
	Use:   "privacy",
	Short: "Grant, revoke, or reset privacy permissions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(privacyCmd).Standalone()
	rootCmd.AddCommand(privacyCmd)
	carapace.Gen(privacyCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionValues("grant", "revoke", "reset"),
		carapace.ActionValues(
			"all",
			"calendar",
			"contacts-limited",
			"contacts",
			"location",
			"location-always",
			"photos-add",
			"photos",
			"media-library",
			"microphone",
			"motion",
			"reminders",
			"siri",
		),
	)
}
