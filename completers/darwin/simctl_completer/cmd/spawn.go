package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var spawnCmd = &cobra.Command{
	Use:   "spawn",
	Short: "Spawn a process by executing a given executable on a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(spawnCmd).Standalone()
	rootCmd.AddCommand(spawnCmd)
	carapace.Gen(spawnCmd).PositionalCompletion(
		simctl.ActionDevicesByState("Booted"),
	)
}
