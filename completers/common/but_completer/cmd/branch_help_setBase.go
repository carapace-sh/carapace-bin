package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_setBaseCmd = &cobra.Command{
	Use:   "set-base",
	Short: "Switch to the GitButler workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_setBaseCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_setBaseCmd)
}
