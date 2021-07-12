package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
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
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.CallbackValue, ".") ||
				strings.HasPrefix(c.CallbackValue, "/") {
				return carapace.ActionFiles()
			}
			return action.ActionPackages(cache_addCmd)
		}),
	)
}
