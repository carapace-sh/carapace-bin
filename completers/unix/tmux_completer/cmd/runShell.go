package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var runShellCmd = &cobra.Command{
	Use:     "run-shell",
	Aliases: []string{"run"},
	Short:   "execute a command without creating a new window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runShellCmd).Standalone()

	runShellCmd.Flags().BoolS("C", "C", false, "execute a tmux command instead of a shell command")
	runShellCmd.Flags().BoolS("E", "E", false, "redirect stderr to stdout")
	runShellCmd.Flags().BoolS("b", "b", false, "run command in background")
	runShellCmd.Flags().StringS("c", "c", "", "specify working directory")
	runShellCmd.Flags().StringS("d", "d", "", "specify delay before starting the command")
	runShellCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(runShellCmd)

	carapace.Gen(runShellCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionDirectories(),
		"t": tmux.ActionPanes(),
	})

	carapace.Gen(runShellCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
