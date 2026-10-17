package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowLsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List windows",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowLsCmd).Standalone()

	windowLsCmd.Flags().BoolP("all", "A", false, "list windows in every session")
	windowLsCmd.Flags().Bool("json", false, "output windows as JSON")
	windowLsCmd.Flags().BoolP("quiet", "q", false, "print IDs only, one per line")
	windowLsCmd.Flags().StringP("session", "s", "", "scope to session SESSION")
	windowLsCmd.Flags().Bool("short", true, "show short IDs (--short=false shows full IDs)")

	windowCmd.AddCommand(windowLsCmd)

	carapace.Gen(windowLsCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(windowLsCmd),
	})
}
