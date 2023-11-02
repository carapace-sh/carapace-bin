package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore paths from another revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_restoreCmd).Standalone()

	helpCmd.AddCommand(help_restoreCmd)
}
