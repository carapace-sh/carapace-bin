package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsCreateContainerCmd = &cobra.Command{
	Use:   "createContainer",
	Short: "Create an empty APFS container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsCreateContainerCmd).Standalone()
	apfsCmd.AddCommand(apfsCreateContainerCmd)
	carapace.Gen(apfsCreateContainerCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
