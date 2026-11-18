package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_compactCmd = &cobra.Command{
	Use:   "compact",
	Short: "Rewrite every record in the compact on-disk encoding and reclaim space",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_compactCmd).Standalone()

	store_compactCmd.Flags().BoolP("help", "h", false, "Print help")
	storeCmd.AddCommand(store_compactCmd)
}
