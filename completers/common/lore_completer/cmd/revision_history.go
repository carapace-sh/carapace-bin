package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_historyCmd = &cobra.Command{
	Use:   "history",
	Short: "List revisions of a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_historyCmd).Standalone()

	revision_historyCmd.Flags().String("branch", "", "Show branch revisions")
	revision_historyCmd.Flags().String("date", "", "Stop when reaching a revision created before this date (Unix timestamp)")
	revision_historyCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_historyCmd.Flags().Bool("oneline", false, "Output each revision on one line only")
	revision_historyCmd.Flags().Bool("only-branch", false, "Stop when reaching a revision on a different branch (includes the branch point revision)")
	revision_historyCmd.Flags().String("revision", "", "Start listing from the specified revision. If not specified, start listing from the current branch latest revision")
	revision_historyCmd.Flag("date").Hidden = true
	revisionCmd.AddCommand(revision_historyCmd)
}
