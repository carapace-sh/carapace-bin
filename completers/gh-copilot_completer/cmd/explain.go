package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain a given input command in natural language",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(explainCmd).Standalone()

	explainCmd.Flags().BoolP("help", "h", false, "help for explain")
	rootCmd.AddCommand(explainCmd)

	carapace.Gen(explainCmd).PositionalCompletion(
		bridge.ActionCarapaceBin().SplitP(),
	)
}
