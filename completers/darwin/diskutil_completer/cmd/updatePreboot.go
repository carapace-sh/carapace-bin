package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsUpdatePrebootCmd = &cobra.Command{
	Use:   "updatePreboot",
	Short: "Update the Preboot volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsUpdatePrebootCmd).Standalone()
	apfsCmd.AddCommand(apfsUpdatePrebootCmd)
	carapace.Gen(apfsUpdatePrebootCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
