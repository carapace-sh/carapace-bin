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

	lock_revokeKeysCmd.Flags().Bool("cosign", false, "cosign the new node-key signature with the current node's tailnet-lock key")
	lock_revokeKeysCmd.Flags().Bool("finish", false, "finish the multi-step revocation process")
	lock_revokeKeysCmd.Flags().String("fork-from", "", "fork the tailnet lock from the given signature")
	lockCmd.AddCommand(lock_revokeKeysCmd)
}
