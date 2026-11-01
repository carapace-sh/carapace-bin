package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selfUpdateCmd = &cobra.Command{
	Use:   "self-update",
	Short: "Update mise itself",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfUpdateCmd).Standalone()

	selfUpdateCmd.Flags().BoolP("force", "f", false, "Update even if already up to date")
	selfUpdateCmd.Flags().BoolP("help", "h", false, "Print help")
	selfUpdateCmd.Flags().String("minimum-release-age", "", "Minimum release age for this update, such as 7d")
	selfUpdateCmd.Flags().Bool("no-plugins", false, "Do not update installed plugins")
	selfUpdateCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")
	rootCmd.AddCommand(selfUpdateCmd)
}
