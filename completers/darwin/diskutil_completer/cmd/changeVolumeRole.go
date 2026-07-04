package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsChangeVolumeRoleCmd = &cobra.Command{
	Use:   "changeVolumeRole",
	Short: "Change role metadata flags on a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsChangeVolumeRoleCmd).Standalone()
	apfsCmd.AddCommand(apfsChangeVolumeRoleCmd)
	carapace.Gen(apfsChangeVolumeRoleCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
