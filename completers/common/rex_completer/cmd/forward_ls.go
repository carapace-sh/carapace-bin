package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var forwardLsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List forwards",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forwardLsCmd).Standalone()

	forwardLsCmd.Flags().BoolP("all", "A", false, "list forwards in every session")
	forwardLsCmd.Flags().Bool("json", false, "output forwards as JSON")
	forwardLsCmd.Flags().BoolP("quiet", "q", false, "print IDs only, one per line")
	forwardLsCmd.Flags().StringP("session", "s", "", "scope to session SESSION")
	forwardLsCmd.Flags().Bool("short", true, "show short IDs (--short=false shows full IDs)")

	forwardCmd.AddCommand(forwardLsCmd)

	carapace.Gen(forwardLsCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(forwardLsCmd),
	})
}
