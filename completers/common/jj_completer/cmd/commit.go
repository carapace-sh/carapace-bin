package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
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

	commitCmd.Flags().String("author", "", "Set author to the provided string")
	commitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	commitCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which changes to include in the first commit")
	commitCmd.Flags().StringSliceP("message", "m", nil, "The change description to use (don't open editor)")
	commitCmd.Flags().Bool("reset-author", false, "Reset the author to the configured user")
	commitCmd.Flags().String("tool", "", "Specify diff editor to be used (implies --interactive)")
	rootCmd.AddCommand(commitCmd)

	carapace.Gen(commitCmd).FlagCompletion(carapace.ActionMap{
		"tool": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	})

	carapace.Gen(commitCmd).PositionalAnyCompletion(
		jj.ActionRevDiffs().FilterArgs(),
	)
}
