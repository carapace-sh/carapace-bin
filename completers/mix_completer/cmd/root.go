package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mix"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
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
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().BoolP("version", "v", false, "show version")

	carapace.Gen(rootCmd).PositionalCompletion(
		mix.ActionMixTasks(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return bridge.ActionCarapaceBin("mix." + c.Args[0]).Shift(1)
		}),
	)
}
