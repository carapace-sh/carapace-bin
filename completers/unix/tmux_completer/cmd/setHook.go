package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var setHookCmd = &cobra.Command{
	Use:   "set-hook",
	Short: "set a hook to a command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setHookCmd).Standalone()

	setHookCmd.Flags().StringS("B", "B", "", "set a subscription to a format")
	setHookCmd.Flags().BoolS("E", "E", false, "fire the user event named hook-name")
	setHookCmd.Flags().BoolS("R", "R", false, "run hook immediately")
	setHookCmd.Flags().BoolS("T", "T", false, "hook is run only when format is true")
	setHookCmd.Flags().BoolS("a", "a", false, "append to hook")
	setHookCmd.Flags().BoolS("g", "g", false, "add hook to global list")
	setHookCmd.Flags().BoolS("p", "p", false, "set pane hooks")
	setHookCmd.Flags().StringS("t", "t", "", "specify target pane")
	setHookCmd.Flags().BoolS("u", "u", false, "unset a hook")
	setHookCmd.Flags().BoolS("w", "w", false, "set window hooks")
	rootCmd.AddCommand(setHookCmd)

	carapace.Gen(setHookCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})

	carapace.Gen(setHookCmd).PositionalCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(setHookCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("tmux").Shift(1),
	)
}
