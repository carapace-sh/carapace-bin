package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var search_semanticCmd = &cobra.Command{
	Use:   "semantic [flags]",
	Short: "Search project code using natural language.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_semanticCmd).Standalone()

	search_semanticCmd.Flags().StringP("directory-path", "d", "", "Restrict search to files under this path (e.g. app/services/).")
	search_semanticCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	search_semanticCmd.Flags().String("knn", "", "Nearest neighbours to retrieve (1–100). Defaults to 64 server-side.")
	search_semanticCmd.Flags().StringP("limit", "l", "", "Maximum number of results (1–100). Defaults to 20 server-side.")
	search_semanticCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	search_semanticCmd.Flags().StringP("query", "q", "", "Natural language search query.")
	search_semanticCmd.MarkFlagRequired("query")
	searchCmd.AddCommand(search_semanticCmd)
}
