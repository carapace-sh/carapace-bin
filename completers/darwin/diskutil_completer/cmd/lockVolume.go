package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsLockVolumeCmd = &cobra.Command{
	Use:   "lockVolume",
	Short: "Unmount and lock an encrypted APFS volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsLockVolumeCmd).Standalone()
	apfsCmd.AddCommand(apfsLockVolumeCmd)
	carapace.Gen(apfsLockVolumeCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
