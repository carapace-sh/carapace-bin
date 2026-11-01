package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_store_immutableCmd = &cobra.Command{
	Use:   "immutable",
	Short: "Operations on the immutable store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_store_immutableCmd).Standalone()

	repository_store_immutableCmd.Flags().BoolP("help", "h", false, "Print help")
	repository_storeCmd.AddCommand(repository_store_immutableCmd)
}
