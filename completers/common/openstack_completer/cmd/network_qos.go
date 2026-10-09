package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_qosCmd = &cobra.Command{
	Use:   "qos",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_qosCmd).Standalone()

	networkCmd.AddCommand(network_qosCmd)
}
