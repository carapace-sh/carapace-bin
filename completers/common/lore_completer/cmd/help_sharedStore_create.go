package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_sharedStore_createCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_sharedStore_createCmd).Standalone()

	help_sharedStoreCmd.AddCommand(help_sharedStore_createCmd)
}
