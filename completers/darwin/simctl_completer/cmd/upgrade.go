package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade a device to a newer runtime",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()
	rootCmd.AddCommand(upgradeCmd)
	carapace.Gen(upgradeCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)
}
