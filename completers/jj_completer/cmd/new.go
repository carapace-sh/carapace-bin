package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [OPTIONS] [REVISIONS]...",
	Short: "Create a new, empty change and edit it in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()

	newCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	newCmd.Flags().BoolP("insert-after", "A", false, "Insert the new change between the target commit(s) and their children")
	newCmd.Flags().BoolP("insert-before", "B", false, "Insert the new change between the target commit(s) and their parents")
	newCmd.Flags().StringSliceP("message", "m", []string{}, "The change description to use")
	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).PositionalAnyCompletion(
		jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
	)
}
