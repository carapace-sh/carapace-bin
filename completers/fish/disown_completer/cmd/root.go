package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "disown",
	Short: "Remove a process from the list of jobs",
	Long:  "https://fishshell.com/docs/current/cmds/disown.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("h", "h", false, "display help")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionJobSpecs(),
	)
}
