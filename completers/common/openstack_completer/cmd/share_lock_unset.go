package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_lock_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Remove resource lock properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_lock_unsetCmd).Standalone()

	share_lock_unsetCmd.Flags().Bool("reason", false, "Unset the lock reason.")
	share_lockCmd.AddCommand(share_lock_unsetCmd)
}
