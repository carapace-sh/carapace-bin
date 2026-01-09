package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/winget"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"update"},
	Short:   "Shows and performs available upgrades",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().Bool("-all", false, "Upgrade all installed packages to latest if available")
	upgradeCmd.Flags().Bool("accept-package-agreements", false, "Accept all license agreements for packages")
	upgradeCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	upgradeCmd.Flags().String("authentication-account", "", "Specify the account to be used for authentication")
	upgradeCmd.Flags().String("authentication-mode", "", "Specify authentication window preference (silent, silentPreferred, or interactive)")
	upgradeCmd.Flags().Bool("force", false, "Direct run the command and continue with non security related issues")
	upgradeCmd.Flags().String("header", "", "Optional Windows-Package-Manager REST source HTTP header")
	upgradeCmd.Flags().String("id", "", "Filter results by id")
	upgradeCmd.Flags().Bool("ignore-local-archive-malware-scan", false, "Ignore the malware scan performed as part of installing an archive type package from local manifest")
	upgradeCmd.Flags().Bool("include-pinned", false, "Upgrade packages even if they have a non-blocking pin")
	upgradeCmd.Flags().Bool("include-unknown", false, "Upgrade packages even if their current version cannot be determined")
	upgradeCmd.Flags().StringP("manifest", "m", "", "The path to the manifest of the package")
	upgradeCmd.Flags().String("name", "", "Filter results by name       ")
	upgradeCmd.Flags().Bool("pinned", false, "Upgrade packages even if they have a non-blocking pin")
	upgradeCmd.Flags().StringP("query", "q", "", "The query used to search for a package")
	upgradeCmd.Flags().BoolP("recurse", "r", false, "Upgrade all installed packages to latest if available")
	upgradeCmd.Flags().Bool("uninstall-previous", false, "Uninstall the previous version of the package during upgrade")
	upgradeCmd.Flags().BoolP("unknown", "u", false, "Upgrade packages even if their current version cannot be determined")
	rootCmd.AddCommand(upgradeCmd)

	// TODO flag completion
	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
		"authentication-account": carapace.ActionValues(),
		"authentication-mode":    winget.ActionAuthenticationModes(),
		"header":                 carapace.ActionValues(),
		"id":                     carapace.ActionValues(),
		"manifest":               carapace.ActionFiles(),
		"name":                   carapace.ActionValues(),
		"query":                  carapace.ActionValues(),
	})
}
