package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_store_rekeyCmd = &cobra.Command{
	Use:   "rekey",
	Short: "Re-encrypt the store with a new key (potential for data loss!)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_store_rekeyCmd).Standalone()

	help_storeCmd.AddCommand(help_store_rekeyCmd)
}
