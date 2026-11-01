package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var object_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset object properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(object_unsetCmd).Standalone()

	object_unsetCmd.Flags().String("property", "", "Property to remove from object (repeat option to remove multiple properties)")
	object_unsetCmd.MarkFlagRequired("property")
	objectCmd.AddCommand(object_unsetCmd)

	carapace.Gen(object_unsetCmd).PositionalCompletion(
		action.ActionContainers(object_unsetCmd),
		action.ActionObjects(object_unsetCmd),
	)
}
