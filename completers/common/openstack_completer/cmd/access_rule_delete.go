package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var access_rule_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete access rule(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_rule_deleteCmd).Standalone()

	access_ruleCmd.AddCommand(access_rule_deleteCmd)
}
