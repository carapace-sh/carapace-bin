package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diagnose_validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate your CLI configuration and connectivity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnose_validateCmd).Standalone()
	diagnoseCmd.AddCommand(diagnose_validateCmd)
}
