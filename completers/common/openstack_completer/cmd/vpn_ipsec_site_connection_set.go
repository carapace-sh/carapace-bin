package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ipsec_site_connection_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set IPsec site connection properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ipsec_site_connection_setCmd).Standalone()

	vpn_ipsec_site_connection_setCmd.Flags().String("description", "", "Description for the connection")
	vpn_ipsec_site_connection_setCmd.Flags().Bool("disable", false, "Disable IPSec site connection")
	vpn_ipsec_site_connection_setCmd.Flags().String("dpd", "", "IPSec Connection Dead Peer Detection attributes.")
	vpn_ipsec_site_connection_setCmd.Flags().Bool("enable", false, "Enable IPSec site connection")
	vpn_ipsec_site_connection_setCmd.Flags().String("initiator", "", "Initiator state")
	vpn_ipsec_site_connection_setCmd.Flags().String("local-endpoint-group", "", "Local endpoint group (name or ID) with subnet(s) for IPsec connection")
	vpn_ipsec_site_connection_setCmd.Flags().String("local-id", "", "An ID to be used instead of the external IP address for a virtual router")
	vpn_ipsec_site_connection_setCmd.Flags().String("mtu", "", "MTU size for the connection")
	vpn_ipsec_site_connection_setCmd.Flags().String("name", "", "Set friendly name for the connection")
	vpn_ipsec_site_connection_setCmd.Flags().String("peer-address", "", "Peer gateway public IPv4/IPv6 address or FQDN")
	vpn_ipsec_site_connection_setCmd.Flags().String("peer-cidr", "", "Remote subnet(s) in CIDR format.")
	vpn_ipsec_site_connection_setCmd.Flags().String("peer-endpoint-group", "", "Peer endpoint group (name or ID) with CIDR(s) for IPSec connection")
	vpn_ipsec_site_connection_setCmd.Flags().String("peer-id", "", "Peer router identity for authentication.")
	vpn_ipsec_site_connectionCmd.AddCommand(vpn_ipsec_site_connection_setCmd)
}
