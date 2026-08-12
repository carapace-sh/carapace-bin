package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var branch_merge_intoCmd = &cobra.Command{
	Use:   "into",
	Short: "Merge into branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_intoCmd).Standalone()

	branch_merge_intoCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_merge_intoCmd.Flags().String("id", "", "ID of the target branch to merge the current branch into")
	branch_merge_intoCmd.Flags().Bool("ignore-links", false, "Merge only the main repository, skipping all linked repositories")
	branch_merge_intoCmd.Flags().String("link", "", "Merge only a specific linked repository at the given mount path")
	branch_mergeCmd.AddCommand(branch_merge_intoCmd)

	carapace.Gen(branch_merge_intoCmd).FlagCompletion(carapace.ActionMap{
		"link": carapace.ActionValues(),
	})

	carapace.Gen(branch_merge_intoCmd).PositionalCompletion(
		action.ActionBranches(branch_merge_intoCmd),
		carapace.ActionValues(), // message
	)
}
