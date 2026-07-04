package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var icloudSyncCmd = &cobra.Command{
	Use:   "icloud_sync",
	Short: "Trigger iCloud sync on a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(icloudSyncCmd).Standalone()
	rootCmd.AddCommand(icloudSyncCmd)
	carapace.Gen(icloudSyncCmd).PositionalCompletion(simctl.ActionDevices())
}
