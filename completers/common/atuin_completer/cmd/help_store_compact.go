package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_store_compactCmd = &cobra.Command{
	Use:   "compact",
	Short: "Rewrite every record in the compact on-disk encoding and reclaim space",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_store_compactCmd).Standalone()

	help_storeCmd.AddCommand(help_store_compactCmd)
}
