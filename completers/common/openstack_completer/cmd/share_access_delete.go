package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_access_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a share access rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_access_deleteCmd).Standalone()

	share_access_deleteCmd.Flags().Bool("unrestrict", false, "Seek access rule deletion despite restrictions.")
	share_access_deleteCmd.Flags().Bool("wait", false, "Wait for share access rule deletion")
	share_accessCmd.AddCommand(share_access_deleteCmd)
}
