package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_firewall_addDropletsCmd = &cobra.Command{
	Use:   "add-droplets",
	Short: "Add Droplets to a cloud firewall",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_firewall_addDropletsCmd).Standalone()
	compute_firewall_addDropletsCmd.Flags().StringSlice("droplet-ids", []string{}, "A comma-separated list of Droplet IDs to place behind the cloud firewall, e.g.: `123,456`")
	compute_firewallCmd.AddCommand(compute_firewall_addDropletsCmd)
}
