package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/winget"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"find"},
	Short:   "Find and show basic info of packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	searchCmd.Flags().String("authentication-account", "", "Specify the account to be used for authentication")
	searchCmd.Flags().String("authentication-mode", "", "Specify authentication window preference (silent, silentPreferred, or interactive)")
	searchCmd.Flags().String("cmd", "", "Filter results by command")
	searchCmd.Flags().String("command", "", "Filter results by command")
	searchCmd.Flags().StringP("count", "n", "", "Show no more than specified number of results")
	searchCmd.Flags().BoolP("exact", "e", false, "Find package using exact match")
	searchCmd.Flags().String("header", "", "Optional Windows-Package-Manager REST source HTTP header")
	searchCmd.Flags().String("id", "", "Filter results by id")
	searchCmd.Flags().String("moniker", "", "Filter results by moniker")
	searchCmd.Flags().String("name", "", "Filter results by name")
	searchCmd.Flags().StringP("query", "q", "", "The query used to search for a package")
	searchCmd.Flags().StringP("source", "s", "", "Find package using the specified source")
	searchCmd.Flags().String("tag", "", "Filter results by tag")
	searchCmd.Flags().Bool("versions", false, "Show available versions of the package")
	rootCmd.AddCommand(searchCmd)

	// TODO flag completion
	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"authentication-account": carapace.ActionValues(),
		"authentication-mode":    winget.ActionAuthenticationModes(),
		"cmd":                    carapace.ActionValues(),
		"command":                carapace.ActionValues(),
		"count":                  carapace.ActionValues(),
		"header":                 carapace.ActionValues(),
		"id":                     carapace.ActionValues(),
		"moniker":                carapace.ActionValues(),
		"name":                   carapace.ActionValues(),
		"query":                  carapace.ActionValues(),
		"source":                 carapace.ActionValues(),
		"tag":                    carapace.ActionValues(),
	})
}
