package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search for packages",
	Aliases: []string{"find", "s", "se"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()
	searchCmd.Flags().String("color", "always", "whether to show color in output")
	searchCmd.Flags().Bool("description", false, "show the description in search")
	searchCmd.Flags().Bool("json", false, "output as json")
	searchCmd.Flags().BoolP("long", "l", false, "Search the registry for packages")
	searchCmd.Flags().Bool("no-description", false, "hide descripton")
	searchCmd.Flags().Bool("offline", false, "force offline mode")
	searchCmd.Flags().BoolP("parseable", "p", false, "output parseable results")
	searchCmd.Flags().Bool("prefer-offline", false, "bypass staleness checks for cached data")
	searchCmd.Flags().Bool("prefer-online", false, "force staleness checks for cached data")
	searchCmd.Flags().String("registry", "", "base URL of the npm registry")
	searchCmd.Flags().String("searchexclude", "", "space separated options that limit the results")
	searchCmd.Flags().Int("searchlimit", 20, "number of items to limit search results to")
	searchCmd.Flags().String("searchopts", "", "space separated search options")

	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValuesDescribed(
			"always", "always show colors",
			"true", "only for file descriptors",
			"false", "never show colors",
		).StyleF(style.ForKeyword),
	})
}
