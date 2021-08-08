package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pub_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cache_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Install a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_addCmd).Standalone()

	cache_addCmd.Flags().Bool("all", false, "Install all matching versions.")
	cache_addCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	cache_addCmd.Flags().StringP("version", "v", "", "Version constraint.")
	cacheCmd.AddCommand(cache_addCmd)

	carapace.Gen(cache_addCmd).FlagCompletion(carapace.ActionMap{
		"version": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionPackageVersions(c.Args[0])
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(cache_addCmd).PositionalCompletion(
		action.ActionPackageSearch(),
	)
}
