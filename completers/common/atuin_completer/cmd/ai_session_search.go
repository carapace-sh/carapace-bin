package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_session_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Full-text search across captured threads, most relevant first.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_session_searchCmd).Standalone()

	ai_session_searchCmd.Flags().String("harness", "", "Only search threads from this harness")
	ai_session_searchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	ai_session_searchCmd.Flags().String("limit", "10", "Maximum threads to return; 0 for unbounded")
	ai_session_searchCmd.Flags().String("style", "auto", "How output is rendered")
	ai_sessionCmd.AddCommand(ai_session_searchCmd)

	carapace.Gen(ai_session_searchCmd).FlagCompletion(carapace.ActionMap{
		"harness": carapace.ActionValues("amp", "claude-code", "codex", "copilot", "grok", "hermes", "kiro", "omp", "opencode", "pi"),
		"style":   carapace.ActionValues("auto", "json", "ndjson", "plain", "pretty"),
	})
}
