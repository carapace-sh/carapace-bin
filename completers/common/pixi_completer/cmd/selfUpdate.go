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

	selfUpdateCmd.Flags().Bool("dry-run", false, "Only show release notes, do not modify the binary")
	selfUpdateCmd.Flags().Bool("force", false, "Force download the desired version when not exactly same with the current. If no desired version, always replace with the latest version")
	selfUpdateCmd.Flags().Bool("no-release-note", false, "Skip printing the release notes")
	selfUpdateCmd.Flags().String("version", "", "The desired version (to downgrade or upgrade to)")
	rootCmd.AddCommand(selfUpdateCmd)
}
