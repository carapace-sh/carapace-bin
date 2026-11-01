package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var service_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete service(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_deleteCmd).Standalone()

	serviceCmd.AddCommand(service_deleteCmd)

	carapace.Gen(service_deleteCmd).PositionalAnyCompletion(action.ActionServices(service_deleteCmd))
}
