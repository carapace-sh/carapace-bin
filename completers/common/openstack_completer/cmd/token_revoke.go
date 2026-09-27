package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var token_revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revoke existing token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_revokeCmd).Standalone()

	tokenCmd.AddCommand(token_revokeCmd)
}
