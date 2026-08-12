package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_help_unresolveCmd = &cobra.Command{
	Use:   "unresolve",
	Short: "Marks the merge unresolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_help_unresolveCmd).Standalone()

	branch_merge_helpCmd.AddCommand(branch_merge_help_unresolveCmd)
}
