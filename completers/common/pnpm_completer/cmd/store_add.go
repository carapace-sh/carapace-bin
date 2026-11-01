package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var store_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Functionally equivalent to pnpm add, except this adds new packages to the store directly without modifying any projects or files outside of the store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_addCmd).Standalone()

	store_addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	storeCmd.AddCommand(store_addCmd)

	carapace.Gen(store_addCmd).PositionalAnyCompletion(
		npm.ActionPackageSearch(""),
	)
}
