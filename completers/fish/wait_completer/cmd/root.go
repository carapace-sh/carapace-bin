package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wait",
	Short: "Wait for jobs to complete",
	Long:  "https://fishshell.com/docs/current/cmds/wait.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("n", "n", false, "return as soon as first completes")
	rootCmd.Flags().Bool("any", false, "return as soon as first completes")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionJobSpecs(),
	)
}
