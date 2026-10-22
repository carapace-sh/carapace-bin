package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var accelerator_arq_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete accelerator request(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_arq_deleteCmd).Standalone()

	accelerator_arqCmd.AddCommand(accelerator_arq_deleteCmd)

	carapace.Gen(accelerator_arq_deleteCmd).PositionalAnyCompletion(action.ActionAcceleratorArqs())
}
