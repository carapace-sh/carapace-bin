package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_tokens_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_tokens_newCmd).Standalone()

	auth_tokensCmd.AddCommand(auth_tokens_newCmd)
}
