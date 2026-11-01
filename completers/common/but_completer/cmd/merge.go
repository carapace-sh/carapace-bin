package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:     "merge",
	Short:   "Merge a branch directly onto the target branch, bypassing review",
	Aliases: []string{"land"},
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "server interactions",
}

func init() {
	carapace.Gen(mergeCmd).Standalone()

	mergeCmd.Flags().Bool("ff-only", false, "Refuse to merge unless the target can be fast-forwarded to the branch tip")
	mergeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	mergeCmd.Flags().Bool("no-ff", false, "Always create a merge commit, even when the branch can be fast-forwarded")
	mergeCmd.Flags().Bool("whole-stack", false, "Merge the entire stack: BRANCH must be the top segment, and the segments below it are published to the target along with it")
	mergeCmd.Flags().Bool("yes", false, "Skip the confirmation prompt")
	rootCmd.AddCommand(mergeCmd)

	carapace.Gen(mergeCmd).PositionalCompletion(
		carapace.Batch(
			but.ActionCliIds(but.CliIdsOpts{Branches: true, Stacks: true}),
			but.ActionLocalBranches(),
		).ToA(),
	)
}
