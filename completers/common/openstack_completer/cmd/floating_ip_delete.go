package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var floating_ip_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete floating IP(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ip_deleteCmd).Standalone()

	floating_ipCmd.AddCommand(floating_ip_deleteCmd)

	carapace.Gen(floating_ip_deleteCmd).PositionalAnyCompletion(action.ActionFloatingIPs())
}
