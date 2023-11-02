package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:     "commit [OPTIONS] [PATHS]...",
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

	carapace.Gen(commitCmd).PositionalAnyCompletion(
		jj.ActionRevDiffs().FilterArgs(),
	)
}
