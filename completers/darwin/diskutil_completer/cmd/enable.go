package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var raidEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a RAID set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(raidEnableCmd).Standalone()
	appleRAIDCmd.AddCommand(raidEnableCmd)
	carapace.Gen(raidEnableCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
