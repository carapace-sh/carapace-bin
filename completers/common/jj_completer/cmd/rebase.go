package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var rebaseCmd = &cobra.Command{
	Use:   "rebase [OPTIONS] --destination <DESTINATION>",
	Short: "Move revisions to different parent(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebaseCmd).Standalone()

	rebaseCmd.Flags().StringSlice("after", nil, "The revision(s) to insert after (can be repeated to create a merge commit)")
	rebaseCmd.Flags().StringSlice("before", nil, "The revision(s) to insert before (can be repeated to create a merge commit)")
	rebaseCmd.Flags().StringSliceP("branch", "b", nil, "Rebase the whole branch relative to destination's ancestors (can be repeated)")
	rebaseCmd.Flags().StringSliceP("destination", "d", nil, "The revision(s) to rebase onto (can be repeated to create a merge commit)")
	rebaseCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rebaseCmd.Flags().StringSliceP("insert-after", "A", nil, "The revision(s) to insert after (can be repeated to create a merge commit)")
	rebaseCmd.Flags().StringSliceP("insert-before", "B", nil, "The revision(s) to insert before (can be repeated to create a merge commit)")
	rebaseCmd.Flags().Bool("keep-divergent", false, "Keep divergent commits while rebasing")
	rebaseCmd.Flags().StringSliceP("revisions", "r", nil, "Rebase the given revisions, rebasing descendants onto this revision's parent(s)")
	rebaseCmd.Flags().Bool("skip-emptied", false, "If true, when rebasing would produce an empty commit, the commit is abandoned. It will not be abandoned if it was already empty before the rebase. Will never skip merge commits with multiple non-empty parents")
	rebaseCmd.Flags().StringSliceP("source", "s", nil, "Rebase specified revision(s) together with their trees of descendants (can be repeated)")
	rootCmd.AddCommand(rebaseCmd)

	carapace.Gen(rebaseCmd).FlagCompletion(carapace.ActionMap{
		"after":         jj.ActionRevs(jj.RevOption{}.Default()),
		"before":        jj.ActionRevs(jj.RevOption{}.Default()),
		"branch":        jj.ActionRevs(jj.RevOption{LocalBookmarks: true, RemoteBookmarks: true, Tags: true}),
		"destination":   jj.ActionRevs(jj.RevOption{}.Default()),
		"insert-after":  jj.ActionRevs(jj.RevOption{}.Default()),
		"insert-before": jj.ActionRevs(jj.RevOption{}.Default()),
		"revisions":     jj.ActionRevs(jj.RevOption{}.Default()),
		"source":        jj.ActionRevs(jj.RevOption{}.Default()),
	})
}
