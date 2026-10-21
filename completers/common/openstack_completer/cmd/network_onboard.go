package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_onboardCmd = &cobra.Command{
	Use:   "onboard",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_onboardCmd).Standalone()

	networkCmd.AddCommand(network_onboardCmd)
}
