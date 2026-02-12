package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kv_help_rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild the KV store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_help_rebuildCmd).Standalone()

	kv_helpCmd.AddCommand(kv_help_rebuildCmd)
}
