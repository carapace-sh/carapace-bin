package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/winget"
	"github.com/spf13/cobra"
)

var pin_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List current pins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pin_listCmd).Standalone()

	pin_listCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	pin_listCmd.Flags().String("authentication-account", "", "Specify the account to be used for authentication")
	pin_listCmd.Flags().String("authentication-mode", "", "Specify authentication window preference")
	pin_listCmd.Flags().String("cmd", "", "Filter results by command")
	pin_listCmd.Flags().String("command", "", "Filter results by command")
	pin_listCmd.Flags().BoolP("exact", "e", false, "Find package using exact match")
	pin_listCmd.Flags().String("header", "", "Optional Windows-Package-Manager REST source HTTP header")
	pin_listCmd.Flags().String("id", "", "Filter results by id")
	pin_listCmd.Flags().String("moniker", "", "Filter results by moniker")
	pin_listCmd.Flags().String("name", "", "Filter results by name")
	pin_listCmd.Flags().StringP("query", "q", "", "The query used to search for a package")
	pin_listCmd.Flags().StringP("source", "s", "", "Find package using the specified source")
	pin_listCmd.Flags().String("tag", "", "Filter results by tag")
	pinCmd.AddCommand(pin_listCmd)

	carapace.Gen(pin_listCmd).FlagCompletion(carapace.ActionMap{
		"authentication-account": carapace.ActionValues(),
		"authentication-mode":    winget.ActionAuthenticationModes(),
		"cmd":                    carapace.ActionValues(),
		"command":                carapace.ActionValues(),
		"header":                 carapace.ActionValues(),
		"id":                     carapace.ActionValues(),
		"moniker":                carapace.ActionValues(),
		"name":                   carapace.ActionValues(),
		"query":                  carapace.ActionValues(),
		"source":                 carapace.ActionValues(),
		"tag":                    carapace.ActionValues(),
	})
}
