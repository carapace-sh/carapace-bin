package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nerdstorageCmd = &cobra.Command{
	Use:   "nerdstorage",
	Short: "Read, write, and delete NerdStorage documents and collections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nerdstorageCmd).Standalone()
	rootCmd.AddCommand(nerdstorageCmd)
}
