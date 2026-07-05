package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var repairVolumeCmd = &cobra.Command{
	Use:   "repairVolume",
	Short: "Repair the filesystem of a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repairVolumeCmd).Standalone()
	rootCmd.AddCommand(repairVolumeCmd)
	carapace.Gen(repairVolumeCmd).PositionalCompletion(fs.ActionBlockDevices())
}
