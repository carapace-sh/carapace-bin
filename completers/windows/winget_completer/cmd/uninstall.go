package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/winget"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls the given package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	uninstallCmd.Flags().Bool("all", false, "Uninstall all versions")
	uninstallCmd.Flags().Bool("all-versions", false, "Uninstall all versions")
	uninstallCmd.Flags().String("authentication-account", "", "Specify the account to be used for authentication")
	uninstallCmd.Flags().String("authentication-mode", "", "Specify authentication window preference")
	uninstallCmd.Flags().BoolP("exact", "e", false, "Find package using exact match")
	uninstallCmd.Flags().Bool("force", false, "Direct run the command and continue with non security related issues")
	uninstallCmd.Flags().String("header", "", "Optional Windows-Package-Manager REST source HTTP header")
	uninstallCmd.Flags().BoolP("interactive", "i", false, "Request interactive installation; user input may be needed")
	uninstallCmd.Flags().StringP("log", "o", "", "Log location (if supported)")
	uninstallCmd.Flags().String("moniker", "", "Filter results by moniker")
	uninstallCmd.Flags().Bool("preserve", false, "Retains all files and directories created by the package (portable)")
	uninstallCmd.Flags().String("product-code", "", "Filters using the product code")
	uninstallCmd.Flags().Bool("purge", false, "Deletes all files and directories in the package directory (portable)")
	uninstallCmd.Flags().String("scope", "", "Select installed package scope filter")
	uninstallCmd.Flags().BoolP("silent", "h", false, "Request silent installation")
	uninstallCmd.Flags().StringP("source", "s", "", "Find package using the specified source")
	uninstallCmd.Flags().StringP("version", "v", "", "The version to act upon")
	rootCmd.AddCommand(uninstallCmd)

	// TODO flag completion
	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"authentication-account": carapace.ActionValues(),
		"authentication-mode":    winget.ActionAuthenticationModes(),
		"header":                 carapace.ActionValues(),
		"log":                    carapace.ActionFiles(),
		"moniker":                carapace.ActionValues(),
		"product-code":           carapace.ActionValues(),
		"scope":                  winget.ActionScopes(),
		"source":                 carapace.ActionValues(),
		"version":                carapace.ActionValues(),
	})
}
