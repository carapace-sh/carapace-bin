package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var shareCmd = &cobra.Command{
	Use:     "share [-l] -a | filesystem",
	Short:   "share a ZFS filesystem",
	GroupID: "mount",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shareCmd).Standalone()

	shareCmd.Flags().BoolS("a", "a", false, "share all available filesystems")
	shareCmd.Flags().BoolS("l", "l", false, "load encryption keys while sharing")

	rootCmd.AddCommand(shareCmd)

	carapace.Gen(shareCmd).PositionalCompletion(
		zfs.ActionFilesystems(),
	)
}
