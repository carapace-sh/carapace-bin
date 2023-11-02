package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Move revisions to different parent(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebaseCmd).Standalone()

	rebaseCmd.Flags().StringSliceP("branch", "b", []string{}, "Rebase the whole branch relative to destination's ancestors (can be repeated)")
	rebaseCmd.Flags().StringSliceP("destination", "d", []string{}, "The revision(s) to rebase onto (can be repeated to create a merge commit)")
	rebaseCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rebaseCmd.Flags().StringP("revision", "r", "", "Rebase only this revision, rebasing descendants onto this revision's parent(s)")
	rebaseCmd.Flags().StringSliceP("source", "s", []string{}, "Rebase specified revision(s) together their tree of descendants (can be repeated)")
	rebaseCmd.MarkFlagRequired("destination")
	rootCmd.AddCommand(rebaseCmd)
}
