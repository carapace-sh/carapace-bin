package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/winget"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"add"},
	Short:   "Installs the selected package",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("accept-package-agreements", false, "Accept all license agreements for packages")
	installCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	installCmd.Flags().Bool("allow-reboot", false, "Allows a reboot if applicable")
	installCmd.Flags().StringP("architecture", "a", "", "Select the architecture")
	installCmd.Flags().String("authentication-account", "", "Specify the account to be used for authentication")
	installCmd.Flags().String("authentication-mode", "", "Specify authentication window preference")
	installCmd.Flags().String("custom", "", "Arguments to be passed on to the installer in addition to the defaults")
	installCmd.Flags().String("dependency-source", "", "Find package dependencies using the specified source")
	installCmd.Flags().BoolP("exact", "e", false, "Find package using exact match")
	installCmd.Flags().Bool("force", false, "Direct run the command and continue with non security related issues")
	installCmd.Flags().String("header", "", "Optional Windows-Package-Manager REST source HTTP header")
	installCmd.Flags().String("id", "", "Filter results by id")
	installCmd.Flags().Bool("ignore-local-archive-malware-scan", false, "Ignore the malware scan performed as part of installing an archive type package from local manifest")
	installCmd.Flags().Bool("ignore-security-hash", false, "Ignore the installer hash check failure")
	installCmd.Flags().String("installer-type", "", "Select the installer type")
	installCmd.Flags().BoolP("interactive", "i", false, "Request interactive installation; user input may be needed")
	installCmd.Flags().String("locale", "", "Locale to use (BCP47 format)")
	installCmd.Flags().StringP("location", "l", "", "Location to install to (if supported)")
	installCmd.Flags().StringP("log", "o", "", "Log location (if supported)")
	installCmd.Flags().StringP("manifest", "m", "", "The path to the manifest of the package")
	installCmd.Flags().String("moniker", "", "Filter results by moniker")
	installCmd.Flags().String("name", "", "Filter results by name")
	installCmd.Flags().Bool("no-upgrade", false, "Skips upgrade if an installed version already exists")
	installCmd.Flags().String("override", "", "Override arguments to be passed on to the installer")
	installCmd.Flags().StringP("query", "q", "", "The query used to search for a package")
	installCmd.Flags().StringP("rename", "r", "", "The value to rename the executable file (portable)")
	installCmd.Flags().String("scope", "", "Select install scope")
	installCmd.Flags().BoolP("silent", "h", false, "Request silent installation")
	installCmd.Flags().Bool("skip-dependencies", false, "Skips processing package dependencies and Windows features")
	installCmd.Flags().StringP("source", "s", "", "Find package using the specified source")
	installCmd.Flags().Bool("uninstall-previous", false, "Uninstall the previous version of the package during upgrade")
	installCmd.Flags().StringP("version", "v", "", "Use the specified version; default is the latest version")
	rootCmd.AddCommand(installCmd)

	// TODO flag completion
	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"architecture":           carapace.ActionValues(),
		"authentication-account": carapace.ActionValues(),
		"authentication-mode":    winget.ActionAuthenticationModes(),
		"custom":                 carapace.ActionValues(),
		"dependency-source":      carapace.ActionValues(),
		"header":                 carapace.ActionValues(),
		"id":                     carapace.ActionValues(),
		"installer-type":         carapace.ActionValues(),
		"locale":                 carapace.ActionValues(),
		"location":               carapace.ActionDirectories(),
		"log":                    carapace.ActionFiles(),
		"manifest":               carapace.ActionFiles(),
		"moniker":                carapace.ActionValues(),
		"name":                   carapace.ActionValues(),
		"override":               carapace.ActionValues(),
		"query":                  carapace.ActionValues(),
		"rename":                 carapace.ActionValues(),
		"scope":                  winget.ActionScopes(),
		"source":                 carapace.ActionValues(),
		"version":                carapace.ActionValues(),
	})

	// TODO positional completion
}
