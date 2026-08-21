package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_type_access_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete access from share type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_type_access_deleteCmd).Standalone()

	share_type_accessCmd.AddCommand(share_type_access_deleteCmd)
}
