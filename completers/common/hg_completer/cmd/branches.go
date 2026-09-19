package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var branchesCmd = &cobra.Command{
	Use:     "branches",
	Short:   "list repository named branches",
	GroupID: groups[group_change_organization].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branchesCmd).Standalone()

	branchesCmd.Flags().BoolP("active", "a", false, "show only branches that have unmerged heads (DEPRECATED)")
	branchesCmd.Flags().BoolP("closed", "c", false, "show normal and closed branches")
	branchesCmd.Flags().StringArrayP("rev", "r", nil, "show branch name(s) of the given rev")
	branchesCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(branchesCmd)

	carapace.Gen(branchesCmd).FlagCompletion(carapace.ActionMap{
		"rev": action.ActionRevisions(),
	})
}
