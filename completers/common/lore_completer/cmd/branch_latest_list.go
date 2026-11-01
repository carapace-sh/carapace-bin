package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lore_completer/cmd/action"
	"github.com/spf13/cobra"
)

var branch_latest_listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_latest_listCmd).Standalone()

	branch_latest_listCmd.Flags().String("branch", "", "Branch to query")
	branch_latest_listCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_latestCmd.AddCommand(branch_latest_listCmd)

	carapace.Gen(branch_latest_listCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(branch_latest_listCmd),
	})

	carapace.Gen(branch_latest_listCmd).PositionalCompletion(
		carapace.ActionValues(), // LIMIT (number)
	)
}
