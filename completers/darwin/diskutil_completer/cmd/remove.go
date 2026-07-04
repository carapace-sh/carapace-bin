package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var raidRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a member from a RAID set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(raidRemoveCmd).Standalone()
	appleRAIDCmd.AddCommand(raidRemoveCmd)
	carapace.Gen(raidRemoveCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
