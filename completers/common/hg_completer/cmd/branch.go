package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hg"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:     "branch",
	Short:   "set or show the current branch name",
	GroupID: groups[group_change_organization].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branchCmd).Standalone()

	branchCmd.Flags().BoolP("clean", "C", false, "reset branch name to parent branch name")
	branchCmd.Flags().BoolP("force", "f", false, "set branch name even if it shadows an existing branch")
	branchCmd.Flags().StringArrayP("rev", "r", nil, "change branches of the given revs (EXPERIMENTAL)")
	rootCmd.AddCommand(branchCmd)

	carapace.Gen(branchCmd).FlagCompletion(carapace.ActionMap{
		"rev": hg.ActionRevisions(),
	})
}
