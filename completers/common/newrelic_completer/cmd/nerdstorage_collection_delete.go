package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nerdstorage_collection_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a NerdStorage collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nerdstorage_collection_deleteCmd).Standalone()
	nerdstorage_collection_deleteCmd.Flags().StringP("collection", "c", "", "the collection name to delete the document from")
	nerdstorage_collection_deleteCmd.Flags().StringP("entityGuid", "e", "", "the entity GUID")
	nerdstorage_collection_deleteCmd.Flags().String("packageId", "", "the external package ID")
	nerdstorage_collection_deleteCmd.Flags().StringP("scope", "s", "USER", "the scope to delete the document from")
	nerdstorage_collectionCmd.AddCommand(nerdstorage_collection_deleteCmd)
}
