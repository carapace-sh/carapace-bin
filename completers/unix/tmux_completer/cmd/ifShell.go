package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var ifShellCmd = &cobra.Command{
	Use:     "if-shell",
	Aliases: []string{"if"},
	Short:   "execute a tmux command if a shell-command succeeded",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ifShellCmd).Standalone()

	ifShellCmd.Flags().BoolS("F", "F", false, "don't execute shell command but use it as a string-value")
	ifShellCmd.Flags().BoolS("b", "b", false, "run shell command in background")
	ifShellCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(ifShellCmd)

	carapace.Gen(ifShellCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})

	carapace.Gen(ifShellCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(ifShellCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("tmux").Shift(1),
	)
}
