package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var csCreateVolumeCmd = &cobra.Command{
	Use:   "createVolume",
	Short: "Create a logical volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(csCreateVolumeCmd).Standalone()
	coreStorageCmd.AddCommand(csCreateVolumeCmd)
	carapace.Gen(csCreateVolumeCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
