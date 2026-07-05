package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var raidCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a logical volume group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(raidCreateCmd).Standalone()
	appleRAIDCmd.AddCommand(raidCreateCmd)
	carapace.Gen(raidCreateCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
