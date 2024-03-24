package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/ollama"
	"github.com/spf13/cobra"
)

var cpCmd = &cobra.Command{
	Use:   "cp SOURCE TARGET",
	Short: "Copy a model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cpCmd).Standalone()

	rootCmd.AddCommand(cpCmd)

	carapace.Gen(cpCmd).PositionalCompletion(
		ollama.ActionModels().MultiParts(":"),
		ollama.ActionModels().MultiParts(":").FilterArgs(),
	)
}
