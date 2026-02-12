package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_kv_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one or more key-value pairs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_kv_deleteCmd).Standalone()

	help_kvCmd.AddCommand(help_kv_deleteCmd)
}
