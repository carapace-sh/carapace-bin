package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_revokeKeysCmd = &cobra.Command{
	Use:   "revoke-keys",
	Short: "Revoke compromised tailnet-lock keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_revokeKeysCmd).Standalone()

	lockCmd.AddCommand(lock_revokeKeysCmd)
}
