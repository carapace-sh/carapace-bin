package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selfUpdateCmd = &cobra.Command{
	Use:   "self-update",
	Short: "Update pixi to the latest version or a specific version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfUpdateCmd).Standalone()

	selfUpdateCmd.Flags().Bool("dry-run", false, "Don't actually update")
	selfUpdateCmd.Flags().Bool("force", false, "Force the update")
	selfUpdateCmd.Flags().Bool("no-release-note", false, "Don't show release notes")
	selfUpdateCmd.Flags().String("version", "", "The desired version")
	rootCmd.AddCommand(selfUpdateCmd)
}
