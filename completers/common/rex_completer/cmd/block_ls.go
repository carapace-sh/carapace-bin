package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var blockLsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List blocks",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockLsCmd).Standalone()

	blockLsCmd.Flags().BoolP("all", "A", false, "list blocks in every session")
	blockLsCmd.Flags().Bool("json", false, "output blocks as JSON")
	blockLsCmd.Flags().BoolP("quiet", "q", false, "print IDs only, one per line")
	blockLsCmd.Flags().StringP("session", "s", "", "scope to session SESSION")
	blockLsCmd.Flags().Bool("short", true, "show short IDs (--short=false shows full IDs)")
	blockLsCmd.Flags().StringP("window", "w", "", "scope to window WINDOW")

	blockCmd.AddCommand(blockLsCmd)

	carapace.Gen(blockLsCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(blockLsCmd),
		"window":  action.ActionWindows(blockLsCmd),
	})
}
