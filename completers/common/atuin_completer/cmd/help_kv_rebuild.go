package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_kv_rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild the KV store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_kv_rebuildCmd).Standalone()

	help_kvCmd.AddCommand(help_kv_rebuildCmd)
}
