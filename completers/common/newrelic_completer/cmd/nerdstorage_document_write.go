package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nerdstorage_document_writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write a NerdStorage document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nerdstorage_document_writeCmd).Standalone()
	nerdstorage_document_writeCmd.Flags().StringP("collection", "c", "", "the collection name to write the document to")
	nerdstorage_document_writeCmd.Flags().StringP("document", "o", "{}", "the document to be written, in JSON format")
	nerdstorage_document_writeCmd.Flags().StringP("documentId", "d", "", "the document ID")
	nerdstorage_document_writeCmd.Flags().StringP("entityGuid", "e", "", "the entity GUID")
	nerdstorage_document_writeCmd.Flags().StringP("packageId", "p", "", "the external package ID")
	nerdstorage_document_writeCmd.Flags().StringP("scope", "s", "USER", "the scope to write the document to")
	nerdstorage_documentCmd.AddCommand(nerdstorage_document_writeCmd)
}
