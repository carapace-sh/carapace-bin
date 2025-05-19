package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var cache_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add the specified packages to local cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_addCmd).Standalone()

	cacheCmd.AddCommand(cache_addCmd)

	carapace.Gen(cache_addCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			action.ActionPackages(cache_addCmd).UnlessF(condition.CompletingPath),
		).ToA(),
	)
}
