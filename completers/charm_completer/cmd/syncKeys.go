package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var syncKeysCmd = &cobra.Command{
	Use:    "sync-keys",
	Short:  "Re-encrypt encrypt keys for all linked public keys",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncKeysCmd).Standalone()

	rootCmd.AddCommand(syncKeysCmd)
}
