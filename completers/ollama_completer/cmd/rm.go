package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/ollama"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm MODEL [MODEL...]",
	Short: "Remove a model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()

	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalAnyCompletion(
		ollama.ActionModels().MultiParts(":").FilterArgs(),
	)
}
