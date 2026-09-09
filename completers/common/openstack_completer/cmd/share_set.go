package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set share properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_setCmd).Standalone()

	share_setCmd.Flags().String("description", "", "New share description.")
	share_setCmd.Flags().String("name", "", "New share name.")
	share_setCmd.Flags().String("property", "", "Set a property to this share (repeat option to set multiple properties)")
	share_setCmd.Flags().String("public", "", "Level of visibility for share.")
	share_setCmd.Flags().String("status", "", "Explicitly update the status of a share (Admin only).")
	share_setCmd.Flags().String("task-state", "", "Indicate which task state to assign the share.")
	shareCmd.AddCommand(share_setCmd)
}
