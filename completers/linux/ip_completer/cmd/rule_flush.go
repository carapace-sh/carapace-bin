package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rule_flushCmd = &cobra.Command{
	Use:   "flush",
	Short: "flush all rules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rule_flushCmd).Standalone()

	ruleCmd.AddCommand(rule_flushCmd)
}
