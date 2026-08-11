package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge two branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_mergeCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_mergeCmd)
}
