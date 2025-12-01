package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var branch_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show commits ahead of base for a specific branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_showCmd).Standalone()

	branch_showCmd.Flags().Bool("ai", false, "Generate AI summary of the branch changes")
	branch_showCmd.Flags().Bool("check", false, "Check if the branch merges cleanly into upstream and identify conflicting commits")
	branch_showCmd.Flags().BoolP("files", "f", false, "Show files modified in each commit with line counts")
	branch_showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	branch_showCmd.Flags().BoolP("review", "r", false, "Fetch and display review information")
	branchCmd.AddCommand(branch_showCmd)

	carapace.Gen(branch_showCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
