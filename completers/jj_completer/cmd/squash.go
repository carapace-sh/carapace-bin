package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var squashCmd = &cobra.Command{
	Use:     "squash",
	Short:   "Move changes from a revision into its parent",
	Aliases: []string{"amend"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(squashCmd).Standalone()

	squashCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	squashCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which parts to squash")
	squashCmd.Flags().StringSliceP("message", "m", []string{}, "The description to use for squashed revision (don't open editor)")
	squashCmd.Flags().StringP("revision", "r", "", "")
	rootCmd.AddCommand(squashCmd)
}
