package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var server_remove_shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Remove a share from a server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_remove_shareCmd).Standalone()

	server_removeCmd.AddCommand(server_remove_shareCmd)

	carapace.Gen(server_remove_shareCmd).PositionalCompletion(
		action.ActionServers(server_remove_shareCmd),
		action.ActionShares(server_remove_shareCmd),
	)
}
