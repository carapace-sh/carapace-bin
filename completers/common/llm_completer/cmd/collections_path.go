package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var collections_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Output the path to the embeddings database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(collections_pathCmd).Standalone()

	collectionsCmd.AddCommand(collections_pathCmd)
}
