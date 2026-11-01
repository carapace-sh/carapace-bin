package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_qos_ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_qos_ruleCmd).Standalone()

	network_qosCmd.AddCommand(network_qos_ruleCmd)
}
