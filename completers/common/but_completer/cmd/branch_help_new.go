package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new branch in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_newCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_newCmd)
}
