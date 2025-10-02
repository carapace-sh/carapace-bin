package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var duplicateCmd = &cobra.Command{
	Use:   "duplicate [OPTIONS] [REVISIONS]...",
	Short: "Create new changes with the same content as existing ones",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duplicateCmd).Standalone()

	duplicateCmd.Flags().StringSlice("after", nil, "The revision(s) to insert after (can be repeated to create a merge commit)")
	duplicateCmd.Flags().StringSlice("before", nil, "The revision(s) to insert before (can be repeated to create a merge commit)")
	duplicateCmd.Flags().StringSliceP("destination", "d", nil, "The revision(s) to duplicate onto (can be repeated to create a merge commit)")
	duplicateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	duplicateCmd.Flags().StringSliceP("insert-after", "A", nil, "The revision(s) to insert after (can be repeated to create a merge commit)")
	duplicateCmd.Flags().StringSliceP("insert-before", "B", nil, "The revision(s) to insert before (can be repeated to create a merge commit)")
	rootCmd.AddCommand(duplicateCmd)

	duplicateCmd.MarkFlagsMutuallyExclusive(
		"after",
		"before",
		"destination",
		"insert-after",
		"insert-before",
	)

	carapace.Gen(duplicateCmd).FlagCompletion(carapace.ActionMap{
		"destination":   jj.ActionRevs(jj.RevOption{}.Default()),
		"insert-after":  jj.ActionRevs(jj.RevOption{}.Default()),
		"insert-before": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(duplicateCmd).PositionalAnyCompletion(
		jj.ActionRevs(jj.RevOption{}.Default()).FilterArgs(),
	)
}
