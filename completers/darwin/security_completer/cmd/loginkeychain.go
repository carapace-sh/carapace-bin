package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var loginKeychainCmd = &cobra.Command{
	Use:   "login-keychain",
	Short: "Display or set the login keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginKeychainCmd).Standalone()
	loginKeychainCmd.Flags().StringP("domain", "d", "", "Use the specified preference domain")
	loginKeychainCmd.Flags().StringP("set", "s", "", "Set the login keychain")
	rootCmd.AddCommand(loginKeychainCmd)
}
