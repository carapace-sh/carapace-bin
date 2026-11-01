package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rule_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore rules table from stdin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rule_restoreCmd).Standalone()

	ruleCmd.AddCommand(rule_restoreCmd)
}
