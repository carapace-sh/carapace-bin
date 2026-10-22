package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var server_dump_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a dump file in server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_dump_createCmd).Standalone()

	server_dumpCmd.AddCommand(server_dump_createCmd)

	carapace.Gen(server_dump_createCmd).PositionalCompletion(action.ActionServers(server_dump_createCmd))
}
