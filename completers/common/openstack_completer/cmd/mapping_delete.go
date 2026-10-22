package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mapping_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete mapping(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mapping_deleteCmd).Standalone()

	mappingCmd.AddCommand(mapping_deleteCmd)

	carapace.Gen(mapping_deleteCmd).PositionalAnyCompletion(action.ActionMappings())
}
