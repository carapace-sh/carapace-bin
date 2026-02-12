package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Manage the atuin data store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storeCmd).Standalone()

	storeCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(storeCmd)
}
