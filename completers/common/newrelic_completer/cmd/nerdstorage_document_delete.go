package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nerdstorage_document_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a NerdStorage document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nerdstorage_document_deleteCmd).Standalone()
	nerdstorage_document_deleteCmd.Flags().StringP("collection", "c", "", "the collection name to delete the document from")
	nerdstorage_document_deleteCmd.Flags().StringP("documentId", "d", "", "the document ID")
	nerdstorage_document_deleteCmd.Flags().StringP("entityGuid", "e", "", "the entity GUID")
	nerdstorage_document_deleteCmd.Flags().StringP("packageId", "p", "", "the external package ID")
	nerdstorage_document_deleteCmd.Flags().StringP("scope", "s", "USER", "the scope to delete the document from")
	nerdstorage_documentCmd.AddCommand(nerdstorage_document_deleteCmd)
}
