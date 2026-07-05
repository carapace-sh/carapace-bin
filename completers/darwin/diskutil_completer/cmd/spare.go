package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var raidSpareCmd = &cobra.Command{
	Use:   "spare",
	Short: "Add a spare to a RAID set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(raidSpareCmd).Standalone()
	appleRAIDCmd.AddCommand(raidSpareCmd)
	carapace.Gen(raidSpareCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
