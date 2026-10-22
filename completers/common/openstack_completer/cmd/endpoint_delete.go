package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var endpoint_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete endpoint(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_deleteCmd).Standalone()

	endpointCmd.AddCommand(endpoint_deleteCmd)

	carapace.Gen(endpoint_deleteCmd).PositionalAnyCompletion(action.ActionEndpoints(endpoint_deleteCmd))
}
