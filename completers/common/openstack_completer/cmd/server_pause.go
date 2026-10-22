package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var server_pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_pauseCmd).Standalone()

	serverCmd.AddCommand(server_pauseCmd)

	carapace.Gen(server_pauseCmd).PositionalCompletion(action.ActionServers(server_pauseCmd))
}
