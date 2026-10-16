package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cache_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Deletes all cache files in mise",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_clearCmd).Standalone()

	cacheCmd.AddCommand(cache_clearCmd)

	carapace.Gen(cache_clearCmd).PositionalAnyCompletion(
		action.ActionTools(),
	)
}
