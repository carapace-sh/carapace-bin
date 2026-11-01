package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var accelerator_device_profile_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete device profile(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_device_profile_deleteCmd).Standalone()

	accelerator_device_profileCmd.AddCommand(accelerator_device_profile_deleteCmd)

	carapace.Gen(accelerator_device_profile_deleteCmd).PositionalAnyCompletion(action.ActionAcceleratorDeviceProfiles(accelerator_device_profile_deleteCmd))
}
