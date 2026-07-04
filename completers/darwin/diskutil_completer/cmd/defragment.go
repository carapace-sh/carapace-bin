package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsDefragmentCmd = &cobra.Command{
	Use:   "defragment",
	Short: "Manage background defragmentation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsDefragmentCmd).Standalone()
	apfsCmd.AddCommand(apfsDefragmentCmd)
	carapace.Gen(apfsDefragmentCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
