package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var container_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset container properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_unsetCmd).Standalone()

	container_unsetCmd.Flags().String("property", "", "Property to remove from container (repeat option to remove multiple properties)")
	container_unsetCmd.MarkFlagRequired("property")
	containerCmd.AddCommand(container_unsetCmd)

	carapace.Gen(container_unsetCmd).PositionalCompletion(action.ActionContainers(container_unsetCmd))
}
