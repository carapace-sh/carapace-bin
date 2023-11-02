package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Manage branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branchCmd).Standalone()

	helpCmd.AddCommand(help_branchCmd)
}