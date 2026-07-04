package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var deleteKeychainCmd = &cobra.Command{
	Use:   "delete-keychain",
	Short: "Delete keychains and remove them from the search list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteKeychainCmd).Standalone()
	rootCmd.AddCommand(deleteKeychainCmd)
	carapace.Gen(deleteKeychainCmd).PositionalAnyCompletion(security.ActionKeychains())
}
