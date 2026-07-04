package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsEraseVolumeCmd = &cobra.Command{
	Use:   "eraseVolume",
	Short: "Erase and format a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsEraseVolumeCmd).Standalone()
	apfsCmd.AddCommand(apfsEraseVolumeCmd)
	carapace.Gen(apfsEraseVolumeCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
