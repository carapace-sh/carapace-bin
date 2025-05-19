package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var store_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds new packages to the store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_addCmd).Standalone()

	storeCmd.AddCommand(store_addCmd)

	carapace.Gen(store_addCmd).PositionalAnyCompletion(
		npm.ActionPackageSearch(""),
	)
}
