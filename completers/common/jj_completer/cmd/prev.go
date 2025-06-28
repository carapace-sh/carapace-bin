package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var prevCmd = &cobra.Command{
	Use:   "prev",
	Short: "Change the working copy revision relative to the parent revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prevCmd).Standalone()

	prevCmd.Flags().Bool("conflict", false, "Jump to the previous conflicted ancestor")
	prevCmd.Flags().BoolP("edit", "e", false, "Edit the parent directly, instead of moving the working-copy commit")
	prevCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	prevCmd.Flags().BoolP("no-edit", "n", false, "The inverse of `--edit`")
	rootCmd.AddCommand(prevCmd)

	prevCmd.MarkFlagsMutuallyExclusive("edit", "no-edit")

	carapace.Gen(prevCmd).PositionalCompletion(
		jj.ActionPrevCommits(100),
	)
}
