package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore working tree and/or index paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_restoreCmd).Standalone()

	helpCmd.AddCommand(help_restoreCmd)
}
