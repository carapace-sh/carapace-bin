package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var unshareCmd = &cobra.Command{
	Use:     "unshare -a | filesystem|mountpoint",
	Short:   "unshare a ZFS filesystem",
	GroupID: "mount",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unshareCmd).Standalone()

	unshareCmd.Flags().BoolS("a", "a", false, "unshare all shared filesystems")

	rootCmd.AddCommand(unshareCmd)

	carapace.Gen(unshareCmd).PositionalCompletion(
		carapace.Batch(
			zfs.ActionFilesystems(),
			carapace.ActionDirectories(),
		).ToA(),
	)
}
