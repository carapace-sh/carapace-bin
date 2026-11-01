package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowRenameCmd = &cobra.Command{
	Use:   "rename [WINDOW] LABEL",
	Short: "Rename a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowRenameCmd).Standalone()

	windowRenameCmd.Flags().StringP("session", "s", "", "scope to session SESSION")

	windowCmd.AddCommand(windowRenameCmd)

	carapace.Gen(windowRenameCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(windowRenameCmd),
	})

	carapace.Gen(windowRenameCmd).PositionalCompletion(
		action.ActionWindows(windowRenameCmd),
	)
}
