package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search packages in crates.io",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("search"),
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().BoolP("help", "h", false, "Print help")
	searchCmd.Flags().String("index", "", "Registry index URL to upload the package to")
	searchCmd.Flags().String("limit", "", "Limit the number of results (default: 10, max: 100)")
	searchCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	searchCmd.Flags().String("registry", "", "Registry to use")
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"registry": action.ActionRegistries(),
	})
}
