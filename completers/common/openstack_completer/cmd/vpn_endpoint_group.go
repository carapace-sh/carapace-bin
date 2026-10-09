package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_endpoint_groupCmd = &cobra.Command{
	Use:   "group",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_endpoint_groupCmd).Standalone()

	vpn_endpointCmd.AddCommand(vpn_endpoint_groupCmd)
}
