package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsResizeContainerCmd = &cobra.Command{
	Use:   "resizeContainer",
	Short: "Resize an APFS container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsResizeContainerCmd).Standalone()
	apfsCmd.AddCommand(apfsResizeContainerCmd)
	carapace.Gen(apfsResizeContainerCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
