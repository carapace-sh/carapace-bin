package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nerdstorage_collectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Read, write, and delete NerdStorage collections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nerdstorage_collectionCmd).Standalone()
	nerdstorageCmd.AddCommand(nerdstorage_collectionCmd)
}
