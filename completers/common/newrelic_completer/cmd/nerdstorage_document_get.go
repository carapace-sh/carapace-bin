package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nerdstorage_document_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a NerdStorage document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nerdstorage_document_getCmd).Standalone()
	nerdstorage_document_getCmd.Flags().StringP("collection", "c", "", "the collection name to get the document from")
	nerdstorage_document_getCmd.Flags().StringP("documentId", "d", "", "the document ID")
	nerdstorage_document_getCmd.Flags().StringP("entityGuid", "e", "", "the entity GUID")
	nerdstorage_document_getCmd.Flags().StringP("packageId", "p", "", "the external package ID")
	nerdstorage_document_getCmd.Flags().StringP("scope", "s", "USER", "the scope to get the document from")
	nerdstorage_documentCmd.AddCommand(nerdstorage_document_getCmd)
}
