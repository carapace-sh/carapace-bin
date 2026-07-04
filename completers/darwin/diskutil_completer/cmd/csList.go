package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var csListCmd = &cobra.Command{
	Use:   "list",
	Short: "List partitions on all disks or a specified disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(csListCmd).Standalone()
	coreStorageCmd.AddCommand(csListCmd)
	carapace.Gen(csListCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
