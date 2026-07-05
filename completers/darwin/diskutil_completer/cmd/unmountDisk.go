package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var unmountDiskCmd = &cobra.Command{
	Use:   "unmountDisk",
	Short: "Unmount all volumes on a disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unmountDiskCmd).Standalone()
	rootCmd.AddCommand(unmountDiskCmd)
	carapace.Gen(unmountDiskCmd).PositionalCompletion(fs.ActionBlockDevices())
}
