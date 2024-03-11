package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var nextCmd = &cobra.Command{
	Use:   "next [OPTIONS] [AMOUNT]",
	Short: "Move the current working copy commit to the next child revision in the repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nextCmd).Standalone()

	nextCmd.Flags().Bool("edit", false, "Instead of creating a new working-copy commit on top of the target commit (like `jj new`), edit the target commit directly (like `jj edit`)")
	nextCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(nextCmd)

	carapace.Gen(nextCmd).PositionalCompletion(
		jj.ActionNextCommits(100),
	)
}
