package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var mountCmd = &cobra.Command{
	Use:     "mount [-flvO] [-o options] -a | filesystem",
	Short:   "mount a ZFS filesystem",
	GroupID: "mount",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mountCmd).Standalone()

	mountCmd.Flags().BoolS("O", "O", false, "perform overlay mount")
	mountCmd.Flags().BoolS("R", "R", false, "mount with all children")
	mountCmd.Flags().BoolS("a", "a", false, "mount all available filesystems")
	mountCmd.Flags().BoolS("f", "f", false, "force mount")
	mountCmd.Flags().BoolP("json", "j", false, "display mounted filesystems in JSON format")
	mountCmd.Flags().BoolS("l", "l", false, "load encryption keys while mounting")
	mountCmd.Flags().StringS("o", "o", "", "temporary mount options")
	mountCmd.Flags().BoolS("v", "v", false, "report mount progress")

	rootCmd.AddCommand(mountCmd)

	carapace.Gen(mountCmd).PositionalCompletion(
		zfs.ActionDatasets(zfs.DatasetOpts{Filesystem: true}),
	)
}
