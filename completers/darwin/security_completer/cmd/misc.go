package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show all commands or usage for a command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var installMdsCmd = &cobra.Command{
	Use:   "install-mds",
	Short: "Install or re-install the MDS database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var leaksCmd = &cobra.Command{
	Use:   "leaks",
	Short: "Run leaks on this process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var setIdentityPreferenceCmd = &cobra.Command{
	Use:   "set-identity-preference",
	Short: "Set the preferred identity for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var getIdentityPreferenceCmd = &cobra.Command{
	Use:   "get-identity-preference",
	Short: "Get the preferred identity for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var executeWithPrivilegesCmd = &cobra.Command{
	Use:   "execute-with-privileges",
	Short: "Execute tool with privileges",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()
	rootCmd.AddCommand(helpCmd)

	carapace.Gen(installMdsCmd).Standalone()
	rootCmd.AddCommand(installMdsCmd)

	carapace.Gen(leaksCmd).Standalone()
	leaksCmd.Flags().Bool("cycles", false, "Use a stricter algorithm")
	leaksCmd.Flags().String("exclude", "", "Ignore leaks called from symbol")
	leaksCmd.Flags().Bool("nocontext", false, "Withhold hex dumps of leaked memory")
	leaksCmd.Flags().Bool("nostacks", false, "Do not show stack traces of leaked memory")
	rootCmd.AddCommand(leaksCmd)

	carapace.Gen(setIdentityPreferenceCmd).Standalone()
	setIdentityPreferenceCmd.Flags().StringP("hash", "Z", "", "Specify identity by SHA-1 hash of certificate")
	setIdentityPreferenceCmd.Flags().StringP("identity", "c", "", "Specify identity by common name of the certificate")
	setIdentityPreferenceCmd.Flags().StringP("key-usage", "u", "", "Specify key usage (optional)")
	setIdentityPreferenceCmd.Flags().BoolP("no-identity", "n", false, "Specify no identity (clears existing preference)")
	setIdentityPreferenceCmd.Flags().StringP("service", "s", "", "Specify service (URL, RFC822 email, DNS host, or other)")
	rootCmd.AddCommand(setIdentityPreferenceCmd)

	carapace.Gen(getIdentityPreferenceCmd).Standalone()
	getIdentityPreferenceCmd.Flags().BoolP("common-name", "c", false, "Print common name of the preferred identity certificate")
	getIdentityPreferenceCmd.Flags().BoolP("hash", "Z", false, "Print SHA-1 hash of the preferred identity certificate")
	getIdentityPreferenceCmd.Flags().StringP("key-usage", "u", "", "Specify key usage (optional)")
	getIdentityPreferenceCmd.Flags().BoolP("pem", "p", false, "Output identity certificate in PEM format")
	getIdentityPreferenceCmd.Flags().StringP("service", "s", "", "Specify service (URL, RFC822 email, DNS host, or other)")
	rootCmd.AddCommand(getIdentityPreferenceCmd)

	carapace.Gen(executeWithPrivilegesCmd).Standalone()
	rootCmd.AddCommand(executeWithPrivilegesCmd)
}
