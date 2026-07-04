package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var imageVolumeCmd = &cobra.Command{
	Use:   "imageVolume",
	Short: "Create and attach a disk image to a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imageVolumeCmd).Standalone()
	rootCmd.AddCommand(imageVolumeCmd)
	carapace.Gen(imageVolumeCmd).PositionalCompletion(fs.ActionBlockDevices())
}
