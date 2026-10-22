package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var local_ip_association_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Local IP association(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(local_ip_association_deleteCmd).Standalone()

	local_ip_associationCmd.AddCommand(local_ip_association_deleteCmd)

	carapace.Gen(local_ip_association_deleteCmd).PositionalCompletion(
		action.ActionLocalIPs(),
		action.ActionPorts(),
	)
}
