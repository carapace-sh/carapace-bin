package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diagnose_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the New Relic Diagnostics binary if necessary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnose_updateCmd).Standalone()
	diagnoseCmd.AddCommand(diagnose_updateCmd)
}
