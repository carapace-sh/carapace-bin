package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpn_endpoint_group_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete endpoint group(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpn_endpoint_group_deleteCmd).Standalone()

	vpn_endpoint_groupCmd.AddCommand(vpn_endpoint_group_deleteCmd)
}
