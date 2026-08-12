package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_help_intoCmd = &cobra.Command{
	Use:   "into",
	Short: "Merge into branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_help_intoCmd).Standalone()

	branch_merge_helpCmd.AddCommand(branch_merge_help_intoCmd)
}
