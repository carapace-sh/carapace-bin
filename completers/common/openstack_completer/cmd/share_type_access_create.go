package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var share_type_access_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add access for share type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_type_access_createCmd).Standalone()

	share_type_accessCmd.AddCommand(share_type_access_createCmd)

	carapace.Gen(share_type_access_createCmd).PositionalCompletion(action.ActionShareTypes())
}
