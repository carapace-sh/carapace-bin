package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var govern_auditCmd = &cobra.Command{
	Use:   "audit <command>",
	Short: "Manage AI agent audit events and sessions. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(govern_auditCmd).Standalone()

	governCmd.AddCommand(govern_auditCmd)
}
