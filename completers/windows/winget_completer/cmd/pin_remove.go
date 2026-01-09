package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/winget"
	"github.com/spf13/cobra"
)

var pin_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a package pin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pin_removeCmd).Standalone()

	pin_removeCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	pin_removeCmd.Flags().String("authentication-account", "", "Specify the account to be used for authentication")
	pin_removeCmd.Flags().String("authentication-mode", "", "Specify authentication window preference (silent, silentPreferred, or interactive)")
	pin_removeCmd.Flags().Bool("cmd", false, "Filter results by command")
	pin_removeCmd.Flags().Bool("command", false, "Filter results by command")
	pin_removeCmd.Flags().BoolP("exact", "e", false, "Find package using exact match")
	pin_removeCmd.Flags().String("header", "", "Optional Windows-Package-Manager REST source HTTP header")
	pin_removeCmd.Flags().String("id", "", "Filter results by id")
	pin_removeCmd.Flags().Bool("installed", false, "Pin a specific installed version")
	pin_removeCmd.Flags().String("moniker", "", "Filter results by moniker")
	pin_removeCmd.Flags().String("name", "", "Filter results by name")
	pin_removeCmd.Flags().StringP("query", "q", "", "The query used to search for a package")
	pin_removeCmd.Flags().StringP("source", "s", "", "Find package using the specified source")
	pin_removeCmd.Flags().String("tag", "", "Filter results by tag")
	pinCmd.AddCommand(pin_removeCmd)

	carapace.Gen(pin_removeCmd).FlagCompletion(carapace.ActionMap{
		"authentication-account": carapace.ActionValues(),
		"authentication-mode":    winget.ActionAuthenticationModes(),
		"header":                 carapace.ActionValues(),
		"id":                     carapace.ActionValues(),
		"moniker":                carapace.ActionValues(),
		"name":                   carapace.ActionValues(),
		"query":                  carapace.ActionValues(),
		"source":                 carapace.ActionValues(),
		"tag":                    carapace.ActionValues(),
	})
}
