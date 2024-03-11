package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var reflog_expireCmd = &cobra.Command{
	Use:   "expire",
	Short: "Prune older reflog entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reflog_expireCmd).Standalone()

	reflog_expireCmd.Flags().Bool("all", false, "Process the reflogs of all references")
	reflog_expireCmd.Flags().BoolP("dry-run", "n", false, "Do not actually prune any entries")
	reflog_expireCmd.Flags().String("expire", "", "Prune entries older than the specified time")
	reflog_expireCmd.Flags().String("expire-unreachable", "", "Prune entries older than <time> that are not reachable from the current tip of the branch")
	reflog_expireCmd.Flags().Bool("rewrite", false, "Adjust a reflogs \"old\" SHA-1 to be equal to the \"new\" SHA-1 field of the entry that now precedes it")
	reflog_expireCmd.Flags().Bool("single-worktree", false, "Limit the processing to reflogs from the current working tree only")
	reflog_expireCmd.Flags().Bool("stale-fix", false, "Prune any reflog entries that point to \"broken commits\"")
	reflog_expireCmd.Flags().Bool("updateref", false, "Update the reference to the value of the top reflog entry")
	reflog_expireCmd.Flags().Bool("verbose", false, "Print extra information on screen.")
	reflogCmd.AddCommand(reflog_expireCmd)

	carapace.Gen(reflog_expireCmd).FlagCompletion(carapace.ActionMap{
		"expire": carapace.ActionValuesDescribed( // TODO custom time completion
			"all", "prunes entries regardless of their age",
			"never", "turns off early pruning of reachable entries",
		),
		"expire-unreachable": carapace.ActionValuesDescribed( // TODO custom time completion
			"all", "prunes entries regardless of their age",
			"never", "turns off early pruning of unreachable entries",
		),
	})

	carapace.Gen(reflog_expireCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{}.Default()).FilterArgs(),
	)
}
