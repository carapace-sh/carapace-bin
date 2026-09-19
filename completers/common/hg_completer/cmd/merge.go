package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:     "merge",
	Short:   "merge another revision into working directory",
	GroupID: groups[group_change_manipulation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeCmd).Standalone()

	mergeCmd.Flags().Bool("abort", false, "abort the ongoing merge")
	mergeCmd.Flags().BoolP("force", "f", false, "force a merge including outstanding changes (DEPRECATED)")
	mergeCmd.Flags().BoolP("preview", "P", false, "review revisions to merge (no merge is performed)")
	mergeCmd.Flags().StringP("rev", "r", "", "revision to merge")
	mergeCmd.Flags().StringP("tool", "t", "", "specify merge tool")
	rootCmd.AddCommand(mergeCmd)

	carapace.Gen(mergeCmd).FlagCompletion(carapace.ActionMap{
		"rev":  action.ActionRevisions(),
		"tool": action.ActionMergeTools(),
	})

	carapace.Gen(mergeCmd).PositionalCompletion(
		action.ActionRevisions(),
	)
}
