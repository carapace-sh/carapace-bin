package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_help_rekeyCmd = &cobra.Command{
	Use:   "rekey",
	Short: "Re-encrypt the store with a new key (potential for data loss!)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_help_rekeyCmd).Standalone()

	store_helpCmd.AddCommand(store_help_rekeyCmd)
}
