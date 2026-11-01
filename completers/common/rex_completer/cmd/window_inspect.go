package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowInspectCmd = &cobra.Command{
	Use:   "inspect [WINDOW]",
	Short: "Inspect a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowInspectCmd).Standalone()

	windowInspectCmd.Flags().Bool("json", false, "output window information as JSON")
	windowInspectCmd.Flags().StringP("session", "s", "", "scope to session SESSION")

	windowCmd.AddCommand(windowInspectCmd)

	carapace.Gen(windowInspectCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(windowInspectCmd),
	})

	carapace.Gen(windowInspectCmd).PositionalCompletion(
		action.ActionWindows(windowInspectCmd),
	)
}
