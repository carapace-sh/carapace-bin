package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var collectionsCmd = &cobra.Command{
	Use:   "collections",
	Short: "View and manage collections of embeddings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(collectionsCmd).Standalone()

	rootCmd.AddCommand(collectionsCmd)
}
