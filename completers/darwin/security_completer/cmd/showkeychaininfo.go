package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var showKeychainInfoCmd = &cobra.Command{
	Use:   "show-keychain-info",
	Short: "Show the settings for a keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showKeychainInfoCmd).Standalone()
	rootCmd.AddCommand(showKeychainInfoCmd)
	carapace.Gen(showKeychainInfoCmd).PositionalAnyCompletion(security.ActionKeychains())
}
