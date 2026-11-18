package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_session_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Capture messages an agent hands over, one JSON object per line on stdin, as capture does its transcript's. Messages already captured are skipped",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_session_pushCmd).Standalone()

	ai_session_pushCmd.Flags().String("harness", "", "The harness the thread is from")
	ai_session_pushCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	ai_session_pushCmd.Flags().String("session", "", "The thread's id, as its harness names it")
	ai_session_pushCmd.MarkFlagRequired("harness")
	ai_session_pushCmd.MarkFlagRequired("session")
	ai_sessionCmd.AddCommand(ai_session_pushCmd)

	carapace.Gen(ai_session_pushCmd).FlagCompletion(carapace.ActionMap{
		"harness": carapace.ActionValues("amp", "claude-code", "codex", "copilot", "grok", "hermes", "kiro", "omp", "opencode", "pi"),
	})
}
