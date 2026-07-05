package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var csUpdateVolumeCmd = &cobra.Command{
	Use:   "updateVolume",
	Short: "Update logical volume properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(csUpdateVolumeCmd).Standalone()
	coreStorageCmd.AddCommand(csUpdateVolumeCmd)
	carapace.Gen(csUpdateVolumeCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
