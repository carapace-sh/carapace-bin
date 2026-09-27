package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_instance_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Forces the deletion of a share instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_instance_deleteCmd).Standalone()

	share_instance_deleteCmd.Flags().Bool("wait", false, "Wait for share instance deletion.")
	share_instanceCmd.AddCommand(share_instance_deleteCmd)
}
