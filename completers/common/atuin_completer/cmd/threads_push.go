package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Capture messages an agent hands over, one JSON object per line on stdin, as capture does its transcript's. Messages already captured are skipped",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_pushCmd).Standalone()

	threads_pushCmd.Flags().String("harness", "", "The harness the thread is from")
	threads_pushCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	threads_pushCmd.Flags().String("session", "", "The thread's id, as its harness names it")
	threads_pushCmd.MarkFlagRequired("harness")
	threads_pushCmd.MarkFlagRequired("session")
	threadsCmd.AddCommand(threads_pushCmd)

	carapace.Gen(threads_pushCmd).FlagCompletion(carapace.ActionMap{
		"harness": carapace.ActionValues("amp", "claude-code", "codex", "copilot", "grok", "hermes", "kiro", "omp", "opencode", "pi"),
	})
}
