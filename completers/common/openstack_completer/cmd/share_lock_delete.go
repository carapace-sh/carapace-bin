package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_lock_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove one or more resource locks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_lock_deleteCmd).Standalone()

	share_lockCmd.AddCommand(share_lock_deleteCmd)
}
