package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var csEncryptVolumeCmd = &cobra.Command{
	Use:   "encryptVolume",
	Short: "Start encryption on a logical volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(csEncryptVolumeCmd).Standalone()
	coreStorageCmd.AddCommand(csEncryptVolumeCmd)
	carapace.Gen(csEncryptVolumeCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
