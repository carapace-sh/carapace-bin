package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/ollama"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push MODEL",
	Short: "Push a model to a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().Bool("insecure", false, "Use an insecure registry")
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).PositionalCompletion(
		ollama.ActionModels().MultiParts(":"),
	)
}
