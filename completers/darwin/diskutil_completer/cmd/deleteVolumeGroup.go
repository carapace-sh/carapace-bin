package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsDeleteVolumeGroupCmd = &cobra.Command{
	Use:   "deleteVolumeGroup",
	Short: "Remove all volumes in an APFS volume group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsDeleteVolumeGroupCmd).Standalone()
	apfsCmd.AddCommand(apfsDeleteVolumeGroupCmd)
	carapace.Gen(apfsDeleteVolumeGroupCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
