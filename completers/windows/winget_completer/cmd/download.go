package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/winget"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Downloads the installer from a given package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downloadCmd).Standalone()

	downloadCmd.Flags().Bool("accept-package-agreements", false, "Accept all license agreements for packages")
	downloadCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	downloadCmd.Flags().StringP("architecture", "a", "", "Select the architecture")
	downloadCmd.Flags().String("authentication-account", "", "Specify the account to be used for authentication")
	downloadCmd.Flags().String("authentication-mode", "", "Specify authentication window preference")
	downloadCmd.Flags().StringP("download-directory", "d", "", "Directory where the installers are downloaded to")
	downloadCmd.Flags().BoolP("exact", "e", false, "Find package using exact match")
	downloadCmd.Flags().String("header", "", "Optional Windows-Package-Manager REST source HTTP header")
	downloadCmd.Flags().String("id", "", "Filter results by id")
	downloadCmd.Flags().Bool("ignore-security-hash", false, "Ignore the installer hash check failure")
	downloadCmd.Flags().String("installer-type", "", "Select the installer type")
	downloadCmd.Flags().String("locale", "", "Locale to use (BCP47 format)")
	downloadCmd.Flags().StringP("manifest", "m", "", "The path to the manifest of the package")
	downloadCmd.Flags().String("moniker", "", "Filter results by moniker")
	downloadCmd.Flags().String("name", "", "Filter results by name")
	downloadCmd.Flags().String("os-version", "", "Target OS version")
	downloadCmd.Flags().String("platform", "", "Select the target platform")
	downloadCmd.Flags().StringP("query", "q", "", "The query used to search for a package")
	downloadCmd.Flags().String("scope", "", "Select install scope")
	downloadCmd.Flags().Bool("skip-dependencies", false, "Skips processing package dependencies and Windows features")
	downloadCmd.Flags().Bool("skip-license", false, "Skips retrieving Microsoft Store package offline license")
	downloadCmd.Flags().Bool("skip-microsoft-store-package-license", false, "Skips retrieving Microsoft Store package offline license")
	downloadCmd.Flags().StringP("source", "s", "", "Find package using the specified source")
	downloadCmd.Flags().StringP("version", "v", "", "Use the specified version; default is the latest version")
	rootCmd.AddCommand(downloadCmd)

	carapace.Gen(downloadCmd).FlagCompletion(carapace.ActionMap{
		"architecture":           carapace.ActionValues(),
		"authentication-account": carapace.ActionValues(),
		"authentication-mode":    winget.ActionAuthenticationModes(),
		"download-directory":     carapace.ActionDirectories(),
		"header":                 carapace.ActionValues(),
		"id":                     carapace.ActionValues(),
		"installer-type":         carapace.ActionValues(),
		"locale":                 carapace.ActionValues(),
		"manifest":               carapace.ActionFiles(),
		"moniker":                carapace.ActionValues(),
		"name":                   carapace.ActionValues(),
		"os-version":             carapace.ActionValues(),
		"platform":               carapace.ActionValues(),
		"query":                  carapace.ActionValues(),
		"scope":                  winget.ActionScopes(),
		"source":                 carapace.ActionValues(),
		"version":                carapace.ActionValues(),
	})
}
