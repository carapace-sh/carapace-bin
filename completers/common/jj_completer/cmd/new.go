package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [OPTIONS] [REVISIONS]...",
	Short: "Create a new, empty change and (by default) edit it in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()

	newCmd.Flags().StringSlice("after", nil, "Insert the new change after the given commit(s)")
	newCmd.Flags().StringSlice("before", nil, "Insert the new change before the given commit(s)")
	newCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	newCmd.Flags().StringSliceP("insert-after", "A", nil, "Insert the new change after the given commit(s)")
	newCmd.Flags().StringSliceP("insert-before", "B", nil, "Insert the new change before the given commit(s)")
	newCmd.Flags().StringSliceP("message", "m", nil, "The change description to use")
	newCmd.Flags().Bool("no-edit", false, "Do not edit the newly created change")
	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).FlagCompletion(carapace.ActionMap{
		"after":         jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
		"before":        jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
		"insert-after":  jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
		"insert-before": jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
	})
	carapace.Gen(newCmd).PositionalAnyCompletion(
		jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
	)
}
