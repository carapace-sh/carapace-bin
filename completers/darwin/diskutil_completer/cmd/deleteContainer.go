package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsDeleteContainerCmd = &cobra.Command{
	Use:   "deleteContainer",
	Short: "Destroy an APFS container and all volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsDeleteContainerCmd).Standalone()
	apfsCmd.AddCommand(apfsDeleteContainerCmd)
	carapace.Gen(apfsDeleteContainerCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
