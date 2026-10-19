package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/windows/scoop_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "show or clear the download cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cacheCmd).Standalone()
	rootCmd.AddCommand(cacheCmd)

	carapace.Gen(cacheCmd).PositionalAnyCompletion(
		action.ActionCachedApps(),
	)
}
