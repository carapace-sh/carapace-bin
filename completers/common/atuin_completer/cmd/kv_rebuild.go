package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kv_rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild the KV store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kv_rebuildCmd).Standalone()

	kv_rebuildCmd.Flags().BoolP("help", "h", false, "Print help")
	kvCmd.AddCommand(kv_rebuildCmd)
}
