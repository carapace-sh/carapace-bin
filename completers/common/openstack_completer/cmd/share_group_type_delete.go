package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_type_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a share group type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_type_deleteCmd).Standalone()

	share_group_typeCmd.AddCommand(share_group_type_deleteCmd)
}
