package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a logical volume group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsCreateCmd).Standalone()
	apfsCmd.AddCommand(apfsCreateCmd)
	carapace.Gen(apfsCreateCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
