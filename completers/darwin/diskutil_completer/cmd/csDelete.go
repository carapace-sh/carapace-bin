package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var csDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a logical volume group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(csDeleteCmd).Standalone()
	coreStorageCmd.AddCommand(csDeleteCmd)
	carapace.Gen(csDeleteCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
