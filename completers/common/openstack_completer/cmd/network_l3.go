package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_l3Cmd = &cobra.Command{
	Use:   "l3",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_l3Cmd).Standalone()

	networkCmd.AddCommand(network_l3Cmd)
}
