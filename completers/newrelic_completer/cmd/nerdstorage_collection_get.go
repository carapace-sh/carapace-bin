package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nerdstorage_collection_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a NerdStorage collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nerdstorage_collection_getCmd).Standalone()
	nerdstorage_collection_getCmd.Flags().StringP("collection", "c", "", "the collection name to get the document from")
	nerdstorage_collection_getCmd.Flags().StringP("entityGuid", "e", "", "the entity GUID")
	nerdstorage_collection_getCmd.Flags().StringP("packageId", "p", "", "the external package ID")
	nerdstorage_collection_getCmd.Flags().StringP("scope", "s", "USER", "the scope to get the document from")
	nerdstorage_collectionCmd.AddCommand(nerdstorage_collection_getCmd)
}
