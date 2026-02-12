package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kv_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys in a namespace, or in all namespaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_help_listCmd).Standalone()

	kv_helpCmd.AddCommand(kv_help_listCmd)
}
