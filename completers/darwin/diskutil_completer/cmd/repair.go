package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var raidRepairCmd = &cobra.Command{
	Use:   "repair",
	Short: "Repair a RAID set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(raidRepairCmd).Standalone()
	appleRAIDCmd.AddCommand(raidRepairCmd)
	carapace.Gen(raidRepairCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
