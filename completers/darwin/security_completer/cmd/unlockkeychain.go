package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var unlockKeychainCmd = &cobra.Command{
	Use:   "unlock-keychain",
	Short: "Unlock a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlockKeychainCmd).Standalone()
	unlockKeychainCmd.Flags().StringP("password", "p", "", "Use password to unlock the keychain")
	rootCmd.AddCommand(unlockKeychainCmd)
	carapace.Gen(unlockKeychainCmd).PositionalAnyCompletion(security.ActionKeychains())
}
