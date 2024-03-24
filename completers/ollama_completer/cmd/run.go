package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/ollama"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run MODEL [PROMPT]",
	Short: "Run a model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().String("format", "", "Response format (e.g. json)")
	runCmd.Flags().Bool("insecure", false, "Use an insecure registry")
	runCmd.Flags().Bool("nowordwrap", false, "Don't wrap words to the next line automatically")
	runCmd.Flags().Bool("verbose", false, "Show timings for response")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		ollama.ActionModels(),
	)
}
