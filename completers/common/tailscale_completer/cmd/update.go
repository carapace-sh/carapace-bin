package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Tailscale to the latest/different version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().Bool("dry-run", false, "print what update would do without doing it")
	updateCmd.Flags().String("track", "", "which track to check for updates: stable, release-candidate, or unstable")
	updateCmd.Flags().String("version", "", "explicit version to update or downgrade to")
	updateCmd.Flags().Bool("yes", false, "update without interactive prompts")
	rootCmd.AddCommand(updateCmd)

	carapace.Gen(updateCmd).FlagCompletion(carapace.ActionMap{
		"track": carapace.ActionValues("stable", "release-candidate", "unstable"),
	})
}
