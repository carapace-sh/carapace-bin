package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/scoop"
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
		scoop.ActionCachedApps(),
	)
}
