package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_subportCmd = &cobra.Command{
	Use:   "subport",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_subportCmd).Standalone()

	networkCmd.AddCommand(network_subportCmd)
}
