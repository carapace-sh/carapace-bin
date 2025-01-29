package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_reloadCmd = &cobra.Command{
	Use:   "reload [options] [CONTAINER...]",
	Short: "Reload firewall rules for one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_reloadCmd).Standalone()

	network_reloadCmd.Flags().BoolP("all", "a", false, "Reload network configuration of all containers")
	network_reloadCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	networkCmd.AddCommand(network_reloadCmd)
}
