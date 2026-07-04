package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var enableOwnershipCmd = &cobra.Command{
	Use:   "enableOwnership",
	Short: "Enable ownership for a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enableOwnershipCmd).Standalone()
	rootCmd.AddCommand(enableOwnershipCmd)
	carapace.Gen(enableOwnershipCmd).PositionalCompletion(fs.ActionBlockDevices())
}
