package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to a different branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_switchCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_switchCmd)
}
