package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pixi"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search a conda package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().StringSliceP("channel", "c", nil, "The channels to consider as a name or a url. Multiple channels can be specified by using this field multiple times")
	searchCmd.Flags().StringP("limit", "l", "", "Limit the number of search results")
	searchCmd.Flags().StringP("platform", "p", "", "The platform to search for, defaults to current platform")
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"platform": pixi.ActionPlatforms(),
	})
}
