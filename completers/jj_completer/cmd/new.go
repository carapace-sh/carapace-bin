package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
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
	newCmd.Flags().StringP("insert-after", "A", "", "Insert the new change between the target commit(s) and their children")
	newCmd.Flags().StringP("insert-before", "B", "", "Insert the new change between the target commit(s) and their parents")
	newCmd.Flags().StringSliceP("message", "m", nil, "The change description to use")
	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).FlagCompletion(carapace.ActionMap{
		"insert-after":  jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
		"insert-before": jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
	})
	carapace.Gen(newCmd).PositionalAnyCompletion(
		jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
	)
}
