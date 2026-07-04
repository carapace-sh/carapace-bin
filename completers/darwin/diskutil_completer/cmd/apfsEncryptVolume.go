package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsEncryptVolumeCmd = &cobra.Command{
	Use:   "encryptVolume",
	Short: "Start encryption on a logical volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsEncryptVolumeCmd).Standalone()
	apfsCmd.AddCommand(apfsEncryptVolumeCmd)
	carapace.Gen(apfsEncryptVolumeCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
