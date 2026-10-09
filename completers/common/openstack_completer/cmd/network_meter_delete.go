package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_meter_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete network meter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_meter_deleteCmd).Standalone()

	network_meterCmd.AddCommand(network_meter_deleteCmd)
}
