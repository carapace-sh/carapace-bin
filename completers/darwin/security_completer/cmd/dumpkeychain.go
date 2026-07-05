package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var dumpKeychainCmd = &cobra.Command{
	Use:   "dump-keychain",
	Short: "Dump the contents of one or more keychains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpKeychainCmd).Standalone()
	dumpKeychainCmd.Flags().BoolP("acl", "a", false, "Dump access control list of items")
	dumpKeychainCmd.Flags().BoolP("data", "d", false, "Dump decrypted data of items")
	dumpKeychainCmd.Flags().BoolP("interactive", "i", false, "Interactive access control list editing mode")
	dumpKeychainCmd.Flags().BoolP("raw", "r", false, "Dump raw encrypted data of items")
	rootCmd.AddCommand(dumpKeychainCmd)
	carapace.Gen(dumpKeychainCmd).PositionalAnyCompletion(security.ActionKeychains())
}
