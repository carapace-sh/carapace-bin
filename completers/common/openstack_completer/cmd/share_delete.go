package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a share",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_deleteCmd).Standalone()

	share_deleteCmd.Flags().Bool("force", false, "Attempt forced removal of share(s), regardless of state (defaults to False)")
	share_deleteCmd.Flags().String("share-group", "", "Optional share group (name or ID) which contains the share")
	share_deleteCmd.Flags().Bool("soft", false, "Soft delete one or more shares.")
	share_deleteCmd.Flags().Bool("wait", false, "Wait for share deletion")
	shareCmd.AddCommand(share_deleteCmd)
}
