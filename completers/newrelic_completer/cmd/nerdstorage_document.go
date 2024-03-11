package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nerdstorage_documentCmd = &cobra.Command{
	Use:   "document",
	Short: "Read, write, and delete NerdStorage documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nerdstorage_documentCmd).Standalone()
	nerdstorageCmd.AddCommand(nerdstorage_documentCmd)
}
