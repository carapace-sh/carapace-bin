package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_store_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Functionally equivalent to pnpm add, except this adds new packages to the store directly without modifying any projects or files outside of the store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_store_addCmd).Standalone()

	help_storeCmd.AddCommand(help_store_addCmd)
}
