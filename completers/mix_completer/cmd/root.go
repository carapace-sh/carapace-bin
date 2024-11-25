package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mix"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mix",
	Short: "A build automation tool for working with applications written in the Elixir programming language",
	Long:  "https://hexdocs.pm/mix/Mix.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		mix.ActionMixTasks(),
	)
}
