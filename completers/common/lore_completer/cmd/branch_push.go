package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var branch_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push commits to remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_pushCmd).Standalone()

	branch_pushCmd.Flags().Bool("fast-forward-merge", false, "Allow the server to fast-forward merge if the target branch head has moved")
	branch_pushCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_pushCmd)

	carapace.Gen(branch_pushCmd).PositionalCompletion(
		action.ActionBranches(branch_pushCmd),
	)
}
