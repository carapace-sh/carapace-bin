package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
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
	splitCmd.Flags().StringSliceS("d", "d", nil, "The revision(s) to rebase the selected changes onto (can be repeated to create a merge commit)")
	splitCmd.Flags().StringSlice("destination", nil, "The revision(s) to rebase the selected changes onto (can be repeated to create a merge commit)")
	splitCmd.Flags().Bool("editor", false, "Open an editor to edit the change description(s)")
	splitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	splitCmd.Flags().StringSliceP("insert-after", "A", nil, "The revision(s) to insert after (can be repeated to create a merge commit)")
	splitCmd.Flags().StringSliceP("insert-before", "B", nil, "The revision(s) to insert before (can be repeated to create a merge commit)")
	splitCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which parts to split")
	splitCmd.Flags().StringSliceP("message", "m", nil, "The change description to use for the selected changes (don't open editor)")
	splitCmd.Flags().StringSliceP("onto", "o", nil, "The revision(s) to rebase the selected changes onto (can be repeated to create a merge commit)")
	splitCmd.Flags().BoolP("parallel", "p", false, "Split the revision into two parallel revisions instead of a parent and child")
	splitCmd.Flags().StringP("revision", "r", "", "The revision to split")
	splitCmd.Flags().String("tool", "", "Specify diff editor to be used (implies --interactive)")
	rootCmd.AddCommand(splitCmd)

	carapace.Gen(splitCmd).FlagCompletion(carapace.ActionMap{
		"after":         jj.ActionRevsets(jj.RevOpts{}.Default()),
		"before":        jj.ActionRevsets(jj.RevOpts{}.Default()),
		"d":             jj.ActionRevsets(jj.RevOpts{}.Default()),
		"destination":   jj.ActionRevsets(jj.RevOpts{}.Default()),
		"insert-after":  jj.ActionRevsets(jj.RevOpts{}.Default()),
		"insert-before": jj.ActionRevsets(jj.RevOpts{}.Default()),
		"onto":          jj.ActionRevsets(jj.RevOpts{}.Default()),
		"revision":      jj.ActionRevsets(jj.RevOpts{}.Default()),
		"tool":          bridge.ActionCarapaceBin().Split(),
	})

	carapace.Gen(splitCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevFiles(splitCmd.Flag("revision").Value.String())
		}),
	)
}
