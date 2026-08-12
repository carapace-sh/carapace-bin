package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_find_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Find revision by metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_find_metadataCmd).Standalone()

	revision_find_metadataCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_findCmd.AddCommand(revision_find_metadataCmd)
}
