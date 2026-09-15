package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var self_updateCmd = &cobra.Command{
	Use:   "self-update",
	Short: "Updates mise itself",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(self_updateCmd).Standalone()

	self_updateCmd.Flags().BoolP("force", "f", false, "Update even if already up to date")
	self_updateCmd.Flags().Bool("no-plugins", false, "Disable auto-updating plugins")
	self_updateCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")
	rootCmd.AddCommand(self_updateCmd)
}
