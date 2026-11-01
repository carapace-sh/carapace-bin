package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade [-v] | upgrade [-V version] -a | pool...",
	Short:   "upgrade pool on-disk version",
	GroupID: "maintenance",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().StringS("V", "V", "", "upgrade to specific version")
	upgradeCmd.Flags().BoolS("a", "a", false, "upgrade all pools")
	upgradeCmd.Flags().BoolS("v", "v", false, "display upgradeable versions")

	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		zfs.ActionPools(),
	)
}
