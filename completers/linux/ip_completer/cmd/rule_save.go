package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rule_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save rules table to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rule_saveCmd).Standalone()

	ruleCmd.AddCommand(rule_saveCmd)
}
