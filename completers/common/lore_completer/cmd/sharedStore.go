package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sharedStoreCmd = &cobra.Command{
	Use:   "shared-store",
	Short: "Manage the shared store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sharedStoreCmd).Standalone()

	sharedStoreCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(sharedStoreCmd)
}
