package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var csDeleteVolumeCmd = &cobra.Command{
	Use:   "deleteVolume",
	Short: "Delete a logical volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(csDeleteVolumeCmd).Standalone()
	coreStorageCmd.AddCommand(csDeleteVolumeCmd)
	carapace.Gen(csDeleteVolumeCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
