package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var lockKeychainCmd = &cobra.Command{
	Use:   "lock-keychain",
	Short: "Lock a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockKeychainCmd).Standalone()
	lockKeychainCmd.Flags().BoolP("all", "a", false, "Lock all keychains")
	rootCmd.AddCommand(lockKeychainCmd)
	carapace.Gen(lockKeychainCmd).PositionalAnyCompletion(security.ActionKeychains())
}
