package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_help_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolves the merge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_help_resolveCmd).Standalone()

	branch_merge_helpCmd.AddCommand(branch_merge_help_resolveCmd)
}
