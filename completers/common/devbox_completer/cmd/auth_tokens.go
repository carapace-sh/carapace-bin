package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auth_tokensCmd = &cobra.Command{
	Use:   "tokens",
	Short: "Manage devbox auth tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_tokensCmd).Standalone()

	authCmd.AddCommand(auth_tokensCmd)
}
