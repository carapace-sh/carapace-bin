package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_optimiseCmd = &cobra.Command{
	Use:   "optimise",
	Short: "replace identical files in the store by hard links",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_optimiseCmd).Standalone()

	storeCmd.AddCommand(store_optimiseCmd)
}
