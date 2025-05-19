package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diagnose_lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Validate your agent config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diagnose_lintCmd).Standalone()
	diagnose_lintCmd.Flags().String("config-file", "", "Path to the config file to be validated.")
	diagnoseCmd.AddCommand(diagnose_lintCmd)

	carapace.Gen(diagnose_lintCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
	})
}
