package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var createKeychainCmd = &cobra.Command{
	Use:   "create-keychain",
	Short: "Create keychains",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createKeychainCmd).Standalone()
	createKeychainCmd.Flags().StringP("password", "p", "", "Use password for the keychains being created")
	createKeychainCmd.Flags().BoolP("prompt-password", "P", false, "Prompt for password using SecurityAgent")
	rootCmd.AddCommand(createKeychainCmd)
	carapace.Gen(createKeychainCmd).PositionalAnyCompletion(security.ActionKeychains())
}
