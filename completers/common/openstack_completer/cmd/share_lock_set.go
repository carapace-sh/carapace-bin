package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var share_lock_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update resource lock properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_lock_setCmd).Standalone()

	share_lock_setCmd.Flags().String("lock-reason", "", "Reason for the resource lock")
	share_lock_setCmd.Flags().String("reason", "", "Reason for the resource lock")
	share_lock_setCmd.Flags().String("resource-action", "", "Resource action to set in the resource lock")
	share_lock_setCmd.Flags().String("resource_action", "", "Resource action to set in the resource lock")
	share_lockCmd.AddCommand(share_lock_setCmd)

	carapace.Gen(share_lock_setCmd).PositionalCompletion(action.ActionShareLocks())
}
