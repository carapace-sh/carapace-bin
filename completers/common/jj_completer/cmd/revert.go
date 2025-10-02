package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Apply the reverse of the given revision(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revertCmd).Standalone()

	revertCmd.Flags().StringSlice("after", nil, "The revision(s) to insert the reverse changes after (can be repeated to create a merge commit)")
	revertCmd.Flags().StringSlice("before", nil, "The revision(s) to insert the reverse changes before (can be repeated to create a merge commit)")
	revertCmd.Flags().StringSliceP("destination", "d", nil, "The revision(s) to apply the reverse changes on top of")
	revertCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	revertCmd.Flags().StringSliceP("insert-after", "A", nil, "The revision(s) to insert the reverse changes after (can be repeated to create a merge commit)")
	revertCmd.Flags().StringSliceP("insert-before", "B", nil, "The revision(s) to insert the reverse changes before (can be repeated to create a merge commit)")
	revertCmd.Flags().StringSliceP("revisions", "r", nil, "The revision(s) to apply the reverse of")
	rootCmd.AddCommand(revertCmd)

	// TODO completion
}
