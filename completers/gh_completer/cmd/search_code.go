package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var search_codeCmd = &cobra.Command{
	Use:   "code <query>",
	Short: "Search within code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(search_codeCmd).Standalone()

	search_codeCmd.Flags().String("extension", "", "Filter on file extension")
	search_codeCmd.Flags().String("filename", "", "Filter on filename")
	search_codeCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	search_codeCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	search_codeCmd.Flags().String("language", "", "Filter results by language")
	search_codeCmd.Flags().StringP("limit", "L", "", "Maximum number of code results to fetch")
	search_codeCmd.Flags().StringSlice("match", nil, "Restrict search to file contents or file path: {file|path}")
	search_codeCmd.Flags().StringSlice("owner", nil, "Filter on owner")
	search_codeCmd.Flags().StringSliceP("repo", "R", nil, "Filter on repository")
	search_codeCmd.Flags().String("size", "", "Filter on size range, in kilobytes")
	search_codeCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	search_codeCmd.Flags().BoolP("web", "w", false, "Open the search query in the web browser")
	searchCmd.AddCommand(search_codeCmd)
}
