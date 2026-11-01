package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push commits to remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_pushCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_pushCmd)
}
