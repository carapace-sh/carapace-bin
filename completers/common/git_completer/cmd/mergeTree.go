package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var mergeTreeCmd = &cobra.Command{
	Use:     "merge-tree",
	Short:   "Perform merge without touching index or working tree",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(mergeTreeCmd).Standalone()

	mergeTreeCmd.Flags().Bool("allow-unrelated-histories", false, "Override common history check check and make the merge proceed anyway.")
	mergeTreeCmd.Flags().String("merge-base", "", "Specify a merge-base for the merge")
	mergeTreeCmd.Flags().Bool("messages", false, "Write any informational messages such as \"Auto-merging <path>\" or CONFLICT notices to the end of stdout")
	mergeTreeCmd.Flags().Bool("name-only", false, "In the Conflicted file info section")
	mergeTreeCmd.Flags().Bool("no-messages", false, "Do not write any informational messages such as \"Auto-merging <path>\" or CONFLICT notices to the end of stdout")
	mergeTreeCmd.Flags().BoolS("z", "z", false, "Do not quote filenames in the <Conflicted file info> section")
	rootCmd.AddCommand(mergeTreeCmd)

	carapace.Gen(mergeTreeCmd).FlagCompletion(carapace.ActionMap{
		"merge-base": git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(mergeTreeCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
