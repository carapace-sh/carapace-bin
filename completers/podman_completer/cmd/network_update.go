package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_updateCmd = &cobra.Command{
	Use:   "update [options] NETWORK",
	Short: "Update an existing podman network",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_updateCmd).Standalone()

	network_updateCmd.Flags().StringSlice("dns-add", []string{}, "add network level nameservers")
	network_updateCmd.Flags().StringSlice("dns-drop", []string{}, "remove network level nameservers")
	networkCmd.AddCommand(network_updateCmd)
}
