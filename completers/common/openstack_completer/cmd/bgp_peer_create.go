package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bgp_peer_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a BGP peer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bgp_peer_createCmd).Standalone()

	bgp_peer_createCmd.Flags().String("auth-type", "", "Authentication algorithm.")
	bgp_peer_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	bgp_peer_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	bgp_peer_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	bgp_peer_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	bgp_peer_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	bgp_peer_createCmd.Flags().String("password", "", "Authentication password")
	bgp_peer_createCmd.Flags().String("peer-ip", "", "Peer IP address")
	bgp_peer_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	bgp_peer_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	bgp_peer_createCmd.Flags().String("project", "", "Owner's project (name or ID)")
	bgp_peer_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	bgp_peer_createCmd.Flags().String("remote-as", "", "Peer AS number.")
	bgp_peer_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	bgp_peer_createCmd.MarkFlagRequired("peer-ip")
	bgp_peer_createCmd.MarkFlagRequired("remote-as")
	bgp_peerCmd.AddCommand(bgp_peer_createCmd)
}
