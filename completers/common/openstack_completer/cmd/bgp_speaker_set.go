package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_speaker_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set BGP speaker properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_speaker_setCmd).Standalone()

	bgp_speaker_setCmd.Flags().Bool("advertise-floating-ip-host-routes", false, "Enable the advertisement of floating IP host routes by the BGP speaker.")
	bgp_speaker_setCmd.Flags().Bool("advertise-tenant-networks", false, "Enable the advertisement of tenant network routes by the BGP speaker.")
	bgp_speaker_setCmd.Flags().String("name", "", "New name for the BGP speaker")
	bgp_speaker_setCmd.Flags().Bool("no-advertise-floating-ip-host-routes", false, "Disable the advertisement of floating IP host routes by the BGP speaker.")
	bgp_speaker_setCmd.Flags().Bool("no-advertise-tenant-networks", false, "Disable the advertisement of tenant network routes by the BGP speaker.")
	bgp_speakerCmd.AddCommand(bgp_speaker_setCmd)
}
