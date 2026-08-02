package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineUpgradeCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "remove any stored offline transaction and delete cached package files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineUpgradeCleanCmd).Standalone()

	offlineUpgradeCmd.AddCommand(offlineUpgradeCleanCmd)
}
