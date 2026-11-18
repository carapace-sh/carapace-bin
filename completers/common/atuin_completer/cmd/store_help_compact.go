package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_help_compactCmd = &cobra.Command{
	Use:   "compact",
	Short: "Rewrite every record in the compact on-disk encoding and reclaim space",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_help_compactCmd).Standalone()

	store_helpCmd.AddCommand(store_help_compactCmd)
}
