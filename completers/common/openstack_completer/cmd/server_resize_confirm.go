package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var server_resize_confirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: "Confirm server resize.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_resize_confirmCmd).Standalone()

	server_resizeCmd.AddCommand(server_resize_confirmCmd)

	carapace.Gen(server_resize_confirmCmd).PositionalCompletion(action.ActionServers())
}
