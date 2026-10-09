package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var pipePaneCmd = &cobra.Command{
	Use:     "pipe-pane",
	Aliases: []string{"pipep"},
	Short:   "pipe output from a pane to a shell command",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipePaneCmd).Standalone()

	pipePaneCmd.Flags().BoolS("I", "I", false, "write stdout from command to the pane as if it were typed")
	pipePaneCmd.Flags().BoolS("O", "O", false, "pipe output from the pane to the command")
	pipePaneCmd.Flags().BoolS("o", "o", false, "only open a pipe if none is currently opened")
	pipePaneCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(pipePaneCmd)

	carapace.Gen(pipePaneCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})

	carapace.Gen(pipePaneCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
