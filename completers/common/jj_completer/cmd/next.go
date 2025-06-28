package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:   "next [OPTIONS] [AMOUNT]",
	Short: "Move the working-copy commit to the child revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nextCmd).Standalone()

	nextCmd.Flags().Bool("conflict", false, "Jump to the next conflicted descendant")
	nextCmd.Flags().BoolP("edit", "e", false, "Instead of creating a new working-copy commit on top of the target commit (like `jj new`), edit the target commit directly (like `jj edit`)")
	nextCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	nextCmd.Flags().BoolP("no-edit", "n", false, "The inverse of `--edit`")
	rootCmd.AddCommand(nextCmd)

	nextCmd.MarkFlagsMutuallyExclusive("edit", "no-edit")

	carapace.Gen(nextCmd).PositionalCompletion(
		jj.ActionNextCommits(100),
	)
}
