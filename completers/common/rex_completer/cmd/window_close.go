package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowCloseCmd = &cobra.Command{
	Use:   "close WINDOW...",
	Short: "Close windows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowCloseCmd).Standalone()

	windowCloseCmd.Flags().StringP("session", "s", "", "scope to session SESSION")

	windowCmd.AddCommand(windowCloseCmd)

	carapace.Gen(windowCloseCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(windowCloseCmd),
	})

	carapace.Gen(windowCloseCmd).PositionalAnyCompletion(
		action.ActionWindows(windowCloseCmd),
	)
}
