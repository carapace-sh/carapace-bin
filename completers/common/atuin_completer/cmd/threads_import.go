package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import every thread the agents have on this machine, including those capture leaves alone: the ones nobody has gone on with since it was turned on",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_importCmd).Standalone()

	threads_importCmd.Flags().String("harness", "", "Only import threads from this harness")
	threads_importCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	threads_importCmd.Flags().String("style", "auto", "How output is rendered")
	threadsCmd.AddCommand(threads_importCmd)

	carapace.Gen(threads_importCmd).FlagCompletion(carapace.ActionMap{
		"harness": carapace.ActionValues("amp", "claude-code", "codex", "copilot", "grok", "hermes", "kiro", "omp", "opencode", "pi"),
		"style":   carapace.ActionValues("auto", "json", "ndjson", "plain", "pretty"),
	})
}
