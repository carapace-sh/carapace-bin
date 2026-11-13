package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var embedModels_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available embedding models",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(embedModels_listCmd).Standalone()

	embedModels_listCmd.Flags().StringP("query", "q", "", "Search for embedding models matching these strings")
	embedModelsCmd.AddCommand(embedModels_listCmd)
}
