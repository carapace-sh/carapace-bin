package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var csResizeVolumeCmd = &cobra.Command{
	Use:   "resizeVolume",
	Short: "Resize a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(csResizeVolumeCmd).Standalone()
	coreStorageCmd.AddCommand(csResizeVolumeCmd)
	carapace.Gen(csResizeVolumeCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
