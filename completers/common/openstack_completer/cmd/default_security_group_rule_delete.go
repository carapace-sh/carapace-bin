package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var default_security_group_rule_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove security group rule(s) from the default security group template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(default_security_group_rule_deleteCmd).Standalone()

	default_security_group_ruleCmd.AddCommand(default_security_group_rule_deleteCmd)
}
