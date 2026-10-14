package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_ipsec_site_connection_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an IPsec site connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_ipsec_site_connection_createCmd).Standalone()

	vpn_ipsec_site_connection_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	vpn_ipsec_site_connection_createCmd.Flags().String("description", "", "Description for the connection")
	vpn_ipsec_site_connection_createCmd.Flags().Bool("disable", false, "Disable IPSec site connection")
	vpn_ipsec_site_connection_createCmd.Flags().String("dpd", "", "IPSec Connection Dead Peer Detection attributes.")
	vpn_ipsec_site_connection_createCmd.Flags().Bool("enable", false, "Enable IPSec site connection")
	vpn_ipsec_site_connection_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	vpn_ipsec_site_connection_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	vpn_ipsec_site_connection_createCmd.Flags().String("ikepolicy", "", "IKE policy associated with this connection (name or ID)")
	vpn_ipsec_site_connection_createCmd.Flags().String("initiator", "", "Initiator state")
	vpn_ipsec_site_connection_createCmd.Flags().String("ipsecpolicy", "", "IPsec policy associated with this connection (name or ID)")
	vpn_ipsec_site_connection_createCmd.Flags().String("local-endpoint-group", "", "Local endpoint group (name or ID) with subnet(s) for IPsec connection")
	vpn_ipsec_site_connection_createCmd.Flags().String("local-id", "", "An ID to be used instead of the external IP address for a virtual router")
	vpn_ipsec_site_connection_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	vpn_ipsec_site_connection_createCmd.Flags().String("mtu", "", "MTU size for the connection")
	vpn_ipsec_site_connection_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	vpn_ipsec_site_connection_createCmd.Flags().String("peer-address", "", "Peer gateway public IPv4/IPv6 address or FQDN")
	vpn_ipsec_site_connection_createCmd.Flags().String("peer-cidr", "", "Remote subnet(s) in CIDR format.")
	vpn_ipsec_site_connection_createCmd.Flags().String("peer-endpoint-group", "", "Peer endpoint group (name or ID) with CIDR(s) for IPSec connection")
	vpn_ipsec_site_connection_createCmd.Flags().String("peer-id", "", "Peer router identity for authentication.")
	vpn_ipsec_site_connection_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	vpn_ipsec_site_connection_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	vpn_ipsec_site_connection_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	vpn_ipsec_site_connection_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	vpn_ipsec_site_connection_createCmd.Flags().String("psk", "", "Pre-shared key string.")
	vpn_ipsec_site_connection_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	vpn_ipsec_site_connection_createCmd.Flags().String("vpnservice", "", "VPN service instance associated with this connection (name or ID)")
	vpn_ipsec_site_connection_createCmd.MarkFlagRequired("ikepolicy")
	vpn_ipsec_site_connection_createCmd.MarkFlagRequired("ipsecpolicy")
	vpn_ipsec_site_connection_createCmd.MarkFlagRequired("peer-address")
	vpn_ipsec_site_connection_createCmd.MarkFlagRequired("peer-id")
	vpn_ipsec_site_connection_createCmd.MarkFlagRequired("psk")
	vpn_ipsec_site_connection_createCmd.MarkFlagRequired("vpnservice")
	vpn_ipsec_site_connectionCmd.AddCommand(vpn_ipsec_site_connection_createCmd)
}
