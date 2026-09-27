package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_type_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a share type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_type_deleteCmd).Standalone()

	share_typeCmd.AddCommand(share_type_deleteCmd)
}
