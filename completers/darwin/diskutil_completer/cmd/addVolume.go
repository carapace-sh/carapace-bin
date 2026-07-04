package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsAddVolumeCmd = &cobra.Command{
	Use:   "addVolume",
	Short: "Add a new APFS volume to a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsAddVolumeCmd).Standalone()
	apfsCmd.AddCommand(apfsAddVolumeCmd)
	carapace.Gen(apfsAddVolumeCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
