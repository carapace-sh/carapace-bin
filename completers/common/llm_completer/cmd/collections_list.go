package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var collections_listCmd = &cobra.Command{
	Use:   "list",
	Short: "View a list of collections",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(collections_listCmd).Standalone()

	collections_listCmd.Flags().StringP("database", "d", "", "Path to embeddings database")
	collections_listCmd.Flags().Bool("json", false, "Output as JSON")
	collectionsCmd.AddCommand(collections_listCmd)
}
