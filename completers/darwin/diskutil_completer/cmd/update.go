package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var raidUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update RAID set properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(raidUpdateCmd).Standalone()
	appleRAIDCmd.AddCommand(raidUpdateCmd)
	carapace.Gen(raidUpdateCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
