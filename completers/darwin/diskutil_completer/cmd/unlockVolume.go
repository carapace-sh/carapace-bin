package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var csUnlockVolumeCmd = &cobra.Command{
	Use:   "unlockVolume",
	Short: "Unlock an encrypted logical volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(csUnlockVolumeCmd).Standalone()
	coreStorageCmd.AddCommand(csUnlockVolumeCmd)
	carapace.Gen(csUnlockVolumeCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
