package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rule_deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "delete a rule",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rule_deleteCmd).Standalone()

	ruleCmd.AddCommand(rule_deleteCmd)
}
