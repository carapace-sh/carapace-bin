package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var sourceFileCmd = &cobra.Command{
	Use:     "source-file",
	Aliases: []string{"source"},
	Short:   "execute tmux commands from a file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sourceFileCmd).Standalone()

	sourceFileCmd.Flags().BoolS("F", "F", false, "expand path as a format")
	sourceFileCmd.Flags().BoolS("n", "n", false, "parse the file but do not execute commands")
	sourceFileCmd.Flags().BoolS("q", "q", false, "don't report error if path doesn't exist")
	sourceFileCmd.Flags().StringS("t", "t", "", "specify target pane")
	sourceFileCmd.Flags().BoolS("v", "v", false, "show the parsed commands and line numbers")
	rootCmd.AddCommand(sourceFileCmd)

	carapace.Gen(sourceFileCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})

	carapace.Gen(sourceFileCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
