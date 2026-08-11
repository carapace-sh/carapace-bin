package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repository_help_store_immutableCmd = &cobra.Command{
	Use:   "immutable",
	Short: "Operations on the immutable store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repository_help_store_immutableCmd).Standalone()

	repository_help_storeCmd.AddCommand(repository_help_store_immutableCmd)
}
