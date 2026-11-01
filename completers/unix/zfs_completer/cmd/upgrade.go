package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade [-v] | upgrade [-r] [-V version] -a|filesystem",
	Short:   "manage on-disk version of ZFS filesystems",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().StringS("V", "V", "", "upgrade to specified version")
	upgradeCmd.Flags().BoolS("a", "a", false, "upgrade all file systems on all imported pools")
	upgradeCmd.Flags().BoolS("r", "r", false, "upgrade file system and all descendent file systems")
	upgradeCmd.Flags().BoolS("v", "v", false, "display currently supported file system versions")

	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		zfs.ActionFilesystems(),
	)
}
