package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var file_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for content in files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_searchCmd).Standalone()

	file_searchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	file_searchCmd.Flags().BoolP("line-number", "n", false, "Prefix each matched line with its 1-based line number within the file")
	file_searchCmd.Flags().Bool("name-only", false, "Print only the paths of files that contain a match, not the matched lines")
	file_searchCmd.Flags().StringP("pattern", "p", "", "The pattern to search for in a single line")
	file_searchCmd.Flags().StringP("revision", "r", "@", "The revision to search files in")
	file_searchCmd.MarkFlagRequired("pattern")
	fileCmd.AddCommand(file_searchCmd)

	carapace.Gen(file_searchCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevsets(jj.RevOpts{}.Default()),
	})
}
