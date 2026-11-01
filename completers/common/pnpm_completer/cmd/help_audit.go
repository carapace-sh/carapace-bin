package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Checks for known security issues with the installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_auditCmd).Standalone()

	helpCmd.AddCommand(help_auditCmd)
}
