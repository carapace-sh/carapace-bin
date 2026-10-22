package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_meterCmd = &cobra.Command{
	Use:   "meter",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_meterCmd).Standalone()

	networkCmd.AddCommand(network_meterCmd)
}
