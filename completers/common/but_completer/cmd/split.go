package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:     "split",
	Short:   "Split a commit in two",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "editing commits",
}

func init() {
	carapace.Gen(splitCmd).Standalone()

	splitCmd.Flags().Bool("allow-merged", false, "Allow targeting branches and commits that are already merged upstream")
	splitCmd.Flags().StringP("file", "F", "", "Read the new commit's message from FILE, or from stdin when FILE is -")
	splitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	splitCmd.Flags().StringSliceP("message", "m", nil, "The message to use for the new commit")
	rootCmd.AddCommand(splitCmd)

	carapace.Gen(splitCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(splitCmd).PositionalAnyCompletion(
		but.ActionCliIds(but.CliIdsOpts{Changes: true}),
	)
}
