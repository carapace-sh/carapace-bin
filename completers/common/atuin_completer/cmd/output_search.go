package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var output_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Full-text search over captured command output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(output_searchCmd).Standalone()

	output_searchCmd.Flags().StringP("context", "C", "", "Show only the matching lines, with this many lines of context on either side; without it, each match's whole output")
	output_searchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	output_searchCmd.Flags().String("limit", "5", "Maximum number of matches to return")
	output_searchCmd.Flags().String("style", "auto", "How matches are rendered")
	outputCmd.AddCommand(output_searchCmd)

	carapace.Gen(output_searchCmd).FlagCompletion(carapace.ActionMap{
		"style": carapace.ActionValues("auto", "json", "ndjson", "plain", "pretty"),
	})
}
