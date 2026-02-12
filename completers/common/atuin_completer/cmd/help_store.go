package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Manage the atuin data store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_storeCmd).Standalone()

	helpCmd.AddCommand(help_storeCmd)
}
