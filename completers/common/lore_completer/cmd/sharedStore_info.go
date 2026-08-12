package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sharedStore_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sharedStore_infoCmd).Standalone()

	sharedStore_infoCmd.Flags().BoolP("help", "h", false, "Print help")
	sharedStoreCmd.AddCommand(sharedStore_infoCmd)
}
