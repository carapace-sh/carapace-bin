package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var defaultKeychainCmd = &cobra.Command{
	Use:   "default-keychain",
	Short: "Display or set the default keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(defaultKeychainCmd).Standalone()
	defaultKeychainCmd.Flags().StringP("domain", "d", "", "Use the specified preference domain")
	defaultKeychainCmd.Flags().StringP("set", "s", "", "Set the default keychain")
	rootCmd.AddCommand(defaultKeychainCmd)
}
