package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var csAddDiskCmd = &cobra.Command{
	Use:   "addDisk",
	Short: "Add a physical disk to a logical volume group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(csAddDiskCmd).Standalone()
	coreStorageCmd.AddCommand(csAddDiskCmd)
	carapace.Gen(csAddDiskCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
