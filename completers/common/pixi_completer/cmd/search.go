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

	searchCmd.Flags().StringP("channel", "c", "", "The channel to search in")
	searchCmd.Flags().StringP("limit", "l", "", "Limit the number of search results")
	searchCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	searchCmd.Flags().StringP("platform", "p", "", "The platform to search for")
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
		"platform":      pixi.ActionPlatforms(),
	})
}
