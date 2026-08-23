package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open app page in App Store.app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(openCmd).Standalone()
	openCmd.Flags().Bool("bundle", false, "Force processing as a bundle ID")
	rootCmd.AddCommand(openCmd)
}
