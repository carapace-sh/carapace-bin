package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show commits ahead of base for a specific branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_showCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_showCmd)
}
