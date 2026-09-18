package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_endpoint_group_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set endpoint group properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_endpoint_group_setCmd).Standalone()

	vpn_endpoint_group_setCmd.Flags().String("description", "", "Description for the endpoint group")
	vpn_endpoint_group_setCmd.Flags().String("name", "", "Set a name for the endpoint group")
	vpn_endpoint_groupCmd.AddCommand(vpn_endpoint_group_setCmd)
}
