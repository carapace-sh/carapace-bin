package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var mergeBaseCmd = &cobra.Command{
	Use:     "merge-base",
	Short:   "Find merge bases or test whether one commit is an ancestor of another",
	GroupID: "history",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeBaseCmd).Standalone()

	mergeBaseCmd.Flags().Bool("all", false, "List every minimal merge base (criss-cross histories)")
	mergeBaseCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	mergeBaseCmd.Flags().Bool("is-ancestor", false, "Exit 0 when `a` is an ancestor of `b`, 1 otherwise; prints nothing in human mode")
	rootCmd.AddCommand(mergeBaseCmd)

	carapace.Gen(mergeBaseCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
