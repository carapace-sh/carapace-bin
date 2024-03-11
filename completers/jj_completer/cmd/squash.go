package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var squashCmd = &cobra.Command{
	Use:     "squash [OPTIONS] [PATHS]...",
	Short:   "Move changes from a revision into its parent",
	Aliases: []string{"amend"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(squashCmd).Standalone()

	squashCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	squashCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which parts to squash")
	squashCmd.Flags().StringSliceP("message", "m", []string{}, "The description to use for squashed revision (don't open editor)")
	squashCmd.Flags().StringP("revision", "r", "@", "")
	rootCmd.AddCommand(squashCmd)

	carapace.Gen(squashCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(squashCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevDiffs(squashCmd.Flag("revision").Value.String()).FilterArgs()
		}),
	)
}
