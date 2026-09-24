package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_meter_rule_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete meter rule(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_meter_rule_deleteCmd).Standalone()

	network_meter_ruleCmd.AddCommand(network_meter_rule_deleteCmd)
}
