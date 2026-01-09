package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/winget"
	"github.com/spf13/cobra"
)

var pin_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new pin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pin_addCmd).Standalone()

	pin_addCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	pin_addCmd.Flags().String("authentication-account", "", "Specify the account to be used for authentication")
	pin_addCmd.Flags().String("authentication-mode", "", "Specify authentication window preference")
	pin_addCmd.Flags().Bool("blocking", false, "Block from upgrading until the pin is removed, preventing override arguments")
	pin_addCmd.Flags().String("cmd", "", "Filter results by command")
	pin_addCmd.Flags().String("command", "", "Filter results by command")
	pin_addCmd.Flags().BoolP("exact", "e", false, "Find package using exact match")
	pin_addCmd.Flags().Bool("force", false, "Direct run the command and continue with non security related issues")
	pin_addCmd.Flags().String("header", "", "Optional Windows-Package-Manager REST source HTTP header")
	pin_addCmd.Flags().String("id", "", "Filter results by id")
	pin_addCmd.Flags().Bool("installed", false, "Pin a specific installed version")
	pin_addCmd.Flags().String("moniker", "", "Filter results by moniker")
	pin_addCmd.Flags().String("name", "", "Filter results by name")
	pin_addCmd.Flags().StringP("query", "q", "", "The query used to search for a package")
	pin_addCmd.Flags().StringP("source", "s", "", "Find package using the specified source")
	pin_addCmd.Flags().String("tag", "", "Filter results by tag")
	pin_addCmd.Flags().StringP("version", "v", "", "Version to which to pin the package")
	pinCmd.AddCommand(pin_addCmd)

	carapace.Gen(pin_addCmd).FlagCompletion(carapace.ActionMap{
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
		"version":                carapace.ActionValues(),
	})
}
