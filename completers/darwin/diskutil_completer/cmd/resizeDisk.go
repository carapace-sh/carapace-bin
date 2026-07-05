package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var csResizeDiskCmd = &cobra.Command{
	Use:   "resizeDisk",
	Short: "Resize a physical disk in a logical volume group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(csResizeDiskCmd).Standalone()
	coreStorageCmd.AddCommand(csResizeDiskCmd)
	carapace.Gen(csResizeDiskCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
