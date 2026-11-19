package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var collections_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the specified collection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(collections_deleteCmd).Standalone()

	collections_deleteCmd.Flags().StringP("database", "d", "", "Path to embeddings database")
	collectionsCmd.AddCommand(collections_deleteCmd)
}
