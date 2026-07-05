package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var raidAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new member to a RAID set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(raidAddCmd).Standalone()
	appleRAIDCmd.AddCommand(raidAddCmd)
	carapace.Gen(raidAddCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
