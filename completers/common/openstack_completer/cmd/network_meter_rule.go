package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_meter_ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_meter_ruleCmd).Standalone()

	network_meterCmd.AddCommand(network_meter_ruleCmd)
}
