package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var csRemoveDiskCmd = &cobra.Command{
	Use:   "removeDisk",
	Short: "Remove a physical disk from a logical volume group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(csRemoveDiskCmd).Standalone()
	coreStorageCmd.AddCommand(csRemoveDiskCmd)
	carapace.Gen(csRemoveDiskCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
