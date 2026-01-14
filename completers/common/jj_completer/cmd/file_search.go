package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
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
	file_searchCmd.Flags().StringP("pattern", "p", "", "The glob pattern to search for")
	file_searchCmd.Flags().StringP("revision", "r", "", "The revision to search files in")
	file_searchCmd.MarkFlagRequired("pattern")
	fileCmd.AddCommand(file_searchCmd)

	carapace.Gen(file_searchCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevSets(jj.RevOption{}.Default()),
	})
}
