package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/windows/scoop_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cache_showCmd = &cobra.Command{
	Use:   "show",
	Short: "show cached files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_showCmd).Standalone()
	cacheCmd.AddCommand(cache_showCmd)

	carapace.Gen(cache_showCmd).PositionalAnyCompletion(
		action.ActionCachedApps(),
	)
}
