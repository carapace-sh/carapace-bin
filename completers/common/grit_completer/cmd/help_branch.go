package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "List branches, or create / delete one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_branchCmd).Standalone()

	helpCmd.AddCommand(help_branchCmd)
}
