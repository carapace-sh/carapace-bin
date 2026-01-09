package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/winget"
	"github.com/spf13/cobra"
)

var repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "Repairs the selected package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repairCmd).Standalone()

	repairCmd.Flags().Bool("accept-package-agreements", false, "Accept all license agreements for packages")
	repairCmd.Flags().Bool("accept-source-agreements", false, "Accept all source agreements during source operations")
	repairCmd.Flags().String("authentication-account", "", "Specify the account to be used for authentication")
	repairCmd.Flags().String("authentication-mode", "", "Specify authentication window preference (silent, silentPreferred, or interactive)")
	repairCmd.Flags().BoolP("exact", "e", false, "Find package using exact match")
	repairCmd.Flags().Bool("force", false, "Direct run the command and continue with non security related issues")
	repairCmd.Flags().String("header", "", "Optional Windows-Package-Manager REST source HTTP header")
	repairCmd.Flags().Bool("ignore-local-archive-malware-scan", false, "Ignore the malware scan performed as part of installing an archive type package from local manifest")
	repairCmd.Flags().Bool("ignore-security-hash", false, "Ignore the installer hash check failure")
	repairCmd.Flags().String("locale", "", "Locale to use (BCP47 format)")
	repairCmd.Flags().StringP("log", "o", "", "Log location (if supported)")
	rootCmd.AddCommand(repairCmd)

	carapace.Gen(repairCmd).FlagCompletion(carapace.ActionMap{
		"authentication-account": carapace.ActionValues(),
		"authentication-mode":    winget.ActionAuthenticationModes(),
		"header":                 carapace.ActionValues(),
		"locale":                 carapace.ActionValues(),
		"log":                    carapace.ActionFiles(),
	})
}
