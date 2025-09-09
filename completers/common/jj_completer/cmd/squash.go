package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var squashCmd = &cobra.Command{
	Use:     "squash [OPTIONS] [PATHS]...",
	Short:   "Move changes from a revision into another revisio",
	Aliases: []string{"amend"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(squashCmd).Standalone()

	squashCmd.Flags().StringArray("after", nil, "Alias for --insert-after")
	squashCmd.Flags().StringArray("before", nil, "Alias for --insert-before")
	squashCmd.Flags().StringSliceP("from", "f", []string{"@"}, "Revision(s) to squash from (default: @)")
	squashCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	squashCmd.Flags().StringSliceP("insert-after", "A", nil, "(Experimental) The revision(s) to insert the new commit after (can be repeated to create a merge commit)")
	squashCmd.Flags().StringSliceP("insert-before", "B", nil, "(Experimental) The revision(s) to insert the new commit before (can be repeated to create a merge commit)")
	squashCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which parts to squash")
	squashCmd.Flags().StringP("into", "t", "@", "Revision to squash into (default: @)")
	squashCmd.Flags().BoolP("keep-emptied", "k", false, "The source revision will not be abandoned")
	squashCmd.Flags().StringSliceP("message", "m", nil, "The description to use for squashed revision (don't open editor)")
	squashCmd.Flags().StringP("revision", "r", "@", "Revision to squash into its parent (default: @). Incompatible with the experimental `-d`/`-A`/`-B` options")
	squashCmd.Flags().String("to", "@", "Revision to squash into (alias for --into)")
	squashCmd.Flags().String("tool", "", "Specify diff editor to be used (implies --interactive)")
	squashCmd.Flags().BoolP("use-destination-message", "u", false, "Use the description of the destination revision and discard the description(s) of the source revision(s)")
	rootCmd.AddCommand(squashCmd)

	carapace.Gen(squashCmd).FlagCompletion(carapace.ActionMap{
		"from":          jj.ActionRevs(jj.RevOption{}.Default()),
		"insert-after":  jj.ActionRevs(jj.RevOption{}.Default()),
		"insert-before": jj.ActionRevs(jj.RevOption{}.Default()),
		"into":          jj.ActionRevs(jj.RevOption{}.Default()),
		"revision":      jj.ActionRevs(jj.RevOption{}.Default()),
		"to":            jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(squashCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevDiffs(squashCmd.Flag("revision").Value.String()).FilterArgs()
		}),
	)
}
