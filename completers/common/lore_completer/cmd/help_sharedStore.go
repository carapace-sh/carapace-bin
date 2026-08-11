package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_sharedStoreCmd = &cobra.Command{
	Use:   "shared-store",
	Short: "Manage the shared store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_sharedStoreCmd).Standalone()

	helpCmd.AddCommand(help_sharedStoreCmd)
}
