package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var setKeychainPasswordCmd = &cobra.Command{
	Use:   "set-keychain-password",
	Short: "Set password for a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setKeychainPasswordCmd).Standalone()
	setKeychainPasswordCmd.Flags().StringP("old-password", "o", "", "Old keychain password")
	setKeychainPasswordCmd.Flags().StringP("password", "p", "", "New keychain password")
	rootCmd.AddCommand(setKeychainPasswordCmd)
	carapace.Gen(setKeychainPasswordCmd).PositionalAnyCompletion(security.ActionKeychains())
}
