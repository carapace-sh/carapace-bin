package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split [OPTIONS] [PATHS]...",
	Short: "Split a revision in two",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splitCmd).Standalone()

	splitCmd.Flags().StringSlice("after", nil, "The revision(s) to insert after (can be repeated to create a merge commit)")
	splitCmd.Flags().StringSlice("before", nil, "The revision(s) to insert before (can be repeated to create a merge commit)")
	splitCmd.Flags().StringSliceP("destination", "d", nil, "The revision(s) to base the new revision onto (can be repeated to create a merge commit)")
	splitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	splitCmd.Flags().StringSliceP("insert-after", "A", nil, "The revision(s) to insert after (can be repeated to create a merge commit)")
	splitCmd.Flags().StringSliceP("insert-before", "B", nil, "The revision(s) to insert before (can be repeated to create a merge commit)")
	splitCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which parts to split")
	splitCmd.Flags().StringSliceP("message", "m", nil, "The change description to use (don't open editor)")
	splitCmd.Flags().BoolP("parallel", "p", false, "Split the revision into two parallel revisions instead of a parent and child")
	splitCmd.Flags().StringP("revision", "r", "", "The revision to split")
	splitCmd.Flags().String("tool", "", "Specify diff editor to be used (implies --interactive)")
	rootCmd.AddCommand(splitCmd)

	// TODO complete more flags
	carapace.Gen(splitCmd).FlagCompletion(carapace.ActionMap{
		"destination": jj.ActionRevs(jj.RevOption{}.Default()),
		"revision":    jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(splitCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevFiles(splitCmd.Flag("revision").Value.String())
		}),
	)
}
