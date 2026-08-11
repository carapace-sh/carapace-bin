package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "List revisions of a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyCmd).Standalone()

	historyCmd.Flags().String("branch", "", "Show branch revisions")
	historyCmd.Flags().String("date", "", "Stop when reaching a revision created before this date (Unix timestamp)")
	historyCmd.Flags().BoolP("help", "h", false, "Print help")
	historyCmd.Flags().Bool("oneline", false, "Output each revision on one line only")
	historyCmd.Flags().Bool("only-branch", false, "Stop when reaching a revision on a different branch (includes the branch point revision)")
	historyCmd.Flags().String("revision", "", "Start listing from the specified revision. If not specified, start listing from the current branch latest revision")
	historyCmd.Flag("date").Hidden = true
	rootCmd.AddCommand(historyCmd)

	carapace.Gen(historyCmd).FlagCompletion(carapace.ActionMap{
		"branch":   action.ActionBranches(historyCmd),
		"revision": action.ActionRevisions(historyCmd),
	})

	carapace.Gen(historyCmd).PositionalCompletion(
		carapace.ActionValues(), // LENGTH (number)
	)
}
