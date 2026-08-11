package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a merge process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_startCmd).Standalone()

	branch_merge_startCmd.Flags().Bool("dry-run", false, "Do a dry run merge start and only report what changes would be done, do not change anything in the file system")
	branch_merge_startCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_merge_startCmd.Flags().String("id", "", "ID of the source branch to merge into the current branch")
	branch_merge_startCmd.Flags().Bool("ignore-links", false, "Merge only the main repository, skipping all linked repositories")
	branch_merge_startCmd.Flags().String("link", "", "Merge only a specific linked repository at the given mount path")
	branch_merge_startCmd.Flags().String("message", "", "Change the message for committing when no conflicts arise from the merge")
	branch_merge_startCmd.Flags().Bool("no-commit", false, "Disable auto commits even if no conflicts arise from the merge")
	branch_mergeCmd.AddCommand(branch_merge_startCmd)
}
