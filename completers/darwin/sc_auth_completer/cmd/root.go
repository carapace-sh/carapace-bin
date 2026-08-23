package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sc_auth",
	Short: "SmartCard authorization setup script",
	Long:  "https://keith.github.io/xcode-manpages/sc_auth.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("C", "C", "", "Country")
	rootCmd.Flags().StringS("E", "E", "", "Email address")
	rootCmd.Flags().StringS("L", "L", "", "Locality")
	rootCmd.Flags().StringS("N", "N", "", "Common name")
	rootCmd.Flags().StringS("O", "O", "", "Organization")
	rootCmd.Flags().StringS("S", "S", "", "State")
	rootCmd.Flags().StringS("U", "U", "", "Organizational unit")
	rootCmd.Flags().StringS("c", "c", "", "Class ID")
	rootCmd.Flags().StringS("d", "d", "", "Domain")
	rootCmd.Flags().BoolS("f", "f", false, "Force")
	rootCmd.Flags().StringS("h", "h", "", "Hash")
	rootCmd.Flags().StringS("k", "k", "", "Keyname")
	rootCmd.Flags().StringS("l", "l", "", "Label")
	rootCmd.Flags().StringS("o", "o", "", "Operation")
	rootCmd.Flags().StringS("p", "p", "", "PIN")
	rootCmd.Flags().StringS("s", "s", "", "Enable/disable/status")
	rootCmd.Flags().StringS("t", "t", "", "Token ID")
	rootCmd.Flags().StringS("u", "u", "", "User")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionFiles(),
		"s": carapace.ActionValues("enable", "disable", "status"),
		"u": os.ActionUsers(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues("pair", "unpair", "pairing_ui", "identities", "list", "changepin", "verifypin", "enable_for_login", "filevault"),
	)
}
