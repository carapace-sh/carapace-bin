package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var disableOwnershipCmd = &cobra.Command{
	Use:   "disableOwnership",
	Short: "Disable ownership for a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(disableOwnershipCmd).Standalone()
	rootCmd.AddCommand(disableOwnershipCmd)
	carapace.Gen(disableOwnershipCmd).PositionalCompletion(fs.ActionBlockDevices())
}
