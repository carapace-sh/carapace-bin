package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowMoveCmd = &cobra.Command{
	Use:   "move [WINDOW] (--before W | --after W | --index N)",
	Short: "Move a window within its session's tab order",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowMoveCmd).Standalone()

	windowMoveCmd.Flags().String("after", "", "move immediately after window W")
	windowMoveCmd.Flags().String("before", "", "move immediately before window W")
	windowMoveCmd.Flags().String("index", "", "move to one-based position N")
	windowMoveCmd.Flags().StringP("session", "s", "", "scope to session SESSION")

	windowCmd.AddCommand(windowMoveCmd)

	carapace.Gen(windowMoveCmd).FlagCompletion(carapace.ActionMap{
		"after":   action.ActionWindows(windowMoveCmd),
		"before":  action.ActionWindows(windowMoveCmd),
		"session": action.ActionSessions(windowMoveCmd),
	})

	carapace.Gen(windowMoveCmd).PositionalAnyCompletion(
		action.ActionWindows(windowMoveCmd),
	)
}
