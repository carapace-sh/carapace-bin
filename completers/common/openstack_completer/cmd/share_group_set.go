package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Explicitly set share group status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_setCmd).Standalone()

	share_group_setCmd.Flags().String("description", "", "Share group description.")
	share_group_setCmd.Flags().String("name", "", "New name for the share group.")
	share_group_setCmd.Flags().String("status", "", "Explicitly update the status of a share group (Admin  only).")
	share_groupCmd.AddCommand(share_group_setCmd)
}
