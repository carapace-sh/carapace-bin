package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_rule_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete security group rule(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_rule_deleteCmd).Standalone()

	security_group_ruleCmd.AddCommand(security_group_rule_deleteCmd)
}
