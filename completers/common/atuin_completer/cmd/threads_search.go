package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Full-text search across captured threads, most relevant first.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_searchCmd).Standalone()

	threads_searchCmd.Flags().String("harness", "", "Only search threads from this harness")
	threads_searchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	threads_searchCmd.Flags().String("limit", "10", "Maximum threads to return; 0 for unbounded")
	threads_searchCmd.Flags().String("style", "auto", "How output is rendered")
	threadsCmd.AddCommand(threads_searchCmd)

	carapace.Gen(threads_searchCmd).FlagCompletion(carapace.ActionMap{
		"harness": carapace.ActionValues("amp", "claude-code", "codex", "copilot", "grok", "hermes", "kiro", "omp", "opencode", "pi"),
		"style":   carapace.ActionValues("auto", "json", "ndjson", "plain", "pretty"),
	})
}
