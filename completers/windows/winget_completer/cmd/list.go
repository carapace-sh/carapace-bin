package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/winget"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "Display installed packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	listCmd.Flags().String("authentication-account", "", "Specify the account to be used for authentication")
	listCmd.Flags().String("authentication-mode", "", "Specify authentication window preference")
	listCmd.Flags().String("cmd", "", "Filter results by command")
	listCmd.Flags().String("command", "", "Filter results by command")
	listCmd.Flags().StringP("count", "n", "", "Show no more than specified number of results (between 1 and 1000)")
	listCmd.Flags().BoolP("exact", "e", false, "Find package using exact match")
	listCmd.Flags().String("header", "", "Optional Windows-Package-Manager REST source HTTP header")
	listCmd.Flags().String("id", "", "Filter results by id")
	listCmd.Flags().Bool("include-pinned", false, "List packages even if they have a pin that prevents upgrade. Can only be used with the --upgrade-available argument")
	listCmd.Flags().Bool("include-unknown", false, "List packages even if their current version cannot be determined. Can only be used with the --upgrade-available argument")
	listCmd.Flags().String("moniker", "", "Filter results by moniker")
	listCmd.Flags().String("name", "", "Filter results by name")
	listCmd.Flags().Bool("pinned", false, "List packages even if they have a pin that prevents upgrade. Can only be used with the --upgrade-available argument")
	listCmd.Flags().StringP("query", "q", "", "The query used to search for a package")
	listCmd.Flags().String("scope", "", "Select installed package scope filter")
	listCmd.Flags().StringP("source", "s", "", "Find package using the specified source")
	listCmd.Flags().String("tag", "", "Filter results by tag")
	listCmd.Flags().BoolP("unknown", "u", false, "List packages even if their current version cannot be determined. Can only be used with the --upgrade-available argument")
	listCmd.Flags().Bool("upgrade-available", false, "Lists only packages which have an upgrade available")
	rootCmd.AddCommand(listCmd)

	// TODO flag completion
	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
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
		"scope":                  winget.ActionScopes(),
		"source":                 carapace.ActionValues(),
		"tag":                    carapace.ActionValues(),
	})
}
