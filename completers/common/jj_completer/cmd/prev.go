package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var prevCmd = &cobra.Command{
	Use:   "prev",
	Short: "Move the working copy commit to the parent of the current revision.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prevCmd).Standalone()

	prevCmd.Flags().Bool("conflict", false, "Jump to the next conflicted descendant")
	prevCmd.Flags().BoolP("edit", "e", false, "Edit the parent directly, instead of moving the working-copy commit")
	prevCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	prevCmd.Flags().Bool("no-edit", false, "Instead of editing the target commit directly (like `jj edit`), create a new working-copy commit on top of the target commit (like `jj new`)")
	rootCmd.AddCommand(prevCmd)

	prevCmd.MarkFlagsMutuallyExclusive("edit", "no-edit")

	carapace.Gen(prevCmd).PositionalCompletion(
		jj.ActionPrevCommits(100),
	)
}
