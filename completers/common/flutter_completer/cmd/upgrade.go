package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade your Flutter development environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().Bool("continue", false, "Continues the upgrade flow starting with the new version of the Flutter tool. Must be used with \"--continue-started-at\".")
	upgradeCmd.Flags().String("continue-started-at", "", "The timestamp at which the top-level upgrade was started in ISO-8601 format.")
	upgradeCmd.Flags().BoolP("force", "f", false, "Force upgrade the flutter branch, potentially discarding local changes to the git repository. Try this if upgrade fails with \"Your flutter checkout has local changes\".")
	upgradeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	upgradeCmd.Flags().Bool("no-force", false, "Force upgrade the flutter branch, potentially discarding local changes to the git repository. Try this if upgrade fails with \"Your flutter checkout has local changes\".")
	upgradeCmd.Flags().Bool("no-verify-only", false, "Verifies for any new version of flutter available, without upgrading to it.")
	upgradeCmd.Flags().Bool("verify-only", false, "Verifies for any new version of flutter available, without upgrading to it.")
	upgradeCmd.Flags().String("working-directory", "", "Override the working directory for the upgrade. Useful for testing upgrade flows without modifying the current flutter repo.")
	rootCmd.AddCommand(upgradeCmd)
	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
		"working-directory": carapace.ActionDirectories(),
	})
}
