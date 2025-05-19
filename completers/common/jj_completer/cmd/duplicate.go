package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var duplicateCmd = &cobra.Command{
	Use:   "duplicate [OPTIONS] [REVISIONS]...",
	Short: "Create a new change with the same content as an existing one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duplicateCmd).Standalone()

	duplicateCmd.Flags().StringArray("after", nil, "Alias for --insert-after")
	duplicateCmd.Flags().StringArray("before", nil, "Alias for --insert-before")
	duplicateCmd.Flags().StringArrayP("destination", "d", []string{"@"}, "The revision(s) to duplicate onto (can be repeated to create a merge commit)")
	duplicateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	duplicateCmd.Flags().StringArrayP("insert-after", "A", nil, "The revision(s) to insert after (can be repeated to create a merge commit)")
	duplicateCmd.Flags().StringArrayP("insert-before", "B", nil, "The revision(s) to insert before (can be repeated to create a merge commit)")
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
