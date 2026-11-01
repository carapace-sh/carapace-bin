package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var luckyCmd = &cobra.Command{
	Use:   "lucky",
	Short: "Install the first app returned from searching the App Store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(luckyCmd).Standalone()
	luckyCmd.Flags().Bool("force", false, "Force reinstall")
	rootCmd.AddCommand(luckyCmd)
}
