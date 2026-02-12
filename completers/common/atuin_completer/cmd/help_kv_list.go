package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_kv_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys in a namespace, or in all namespaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_kv_listCmd).Standalone()

	help_kvCmd.AddCommand(help_kv_listCmd)
}
