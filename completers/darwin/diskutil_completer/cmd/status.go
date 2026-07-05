package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var raidStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status of a RAID set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(raidStatusCmd).Standalone()
	appleRAIDCmd.AddCommand(raidStatusCmd)
	carapace.Gen(raidStatusCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
