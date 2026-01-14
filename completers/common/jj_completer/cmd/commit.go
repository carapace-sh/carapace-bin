package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Update the description and create a new change on top [default alias: ci]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitCmd).Standalone()

	commitCmd.Flags().Bool("editor", false, "Open an editor to edit the change description")
	commitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	commitCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which changes to include in the current commit")
	commitCmd.Flags().StringSliceP("message", "m", nil, "The change description to use (don't open editor)")
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
