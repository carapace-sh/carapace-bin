package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kv_help_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one or more key-value pairs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_help_deleteCmd).Standalone()

	kv_helpCmd.AddCommand(kv_help_deleteCmd)
}
