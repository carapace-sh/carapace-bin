package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_store_immutableCmd = &cobra.Command{
	Use:   "immutable",
	Short: "Operations on the immutable store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_store_immutableCmd).Standalone()

	help_repository_storeCmd.AddCommand(help_repository_store_immutableCmd)
}
