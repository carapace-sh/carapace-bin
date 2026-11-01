package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var unmountCmd = &cobra.Command{
	Use:     "unmount [-fu] -a | filesystem|mountpoint",
	Short:   "unmount a ZFS filesystem",
	Aliases: []string{"umount"},
	GroupID: "mount",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unmountCmd).Standalone()

	unmountCmd.Flags().BoolS("a", "a", false, "unmount all available filesystems")
	unmountCmd.Flags().BoolS("f", "f", false, "force unmount")
	unmountCmd.Flags().BoolS("u", "u", false, "unload encryption keys")

	rootCmd.AddCommand(unmountCmd)

	carapace.Gen(unmountCmd).PositionalCompletion(
		carapace.Batch(
			zfs.ActionMountedFilesystems(),
			carapace.ActionDirectories(),
		).ToA(),
	)
}
