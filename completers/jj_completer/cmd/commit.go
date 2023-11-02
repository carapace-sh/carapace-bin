package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:     "commit",
	Short:   "Update the description and create a new change on top",
	Aliases: []string{"ci"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitCmd).Standalone()

	commitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	commitCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which changes to include in the first commit")
	commitCmd.Flags().StringSliceP("message", "m", []string{}, "The change description to use (don't open editor)")
	rootCmd.AddCommand(commitCmd)
}
