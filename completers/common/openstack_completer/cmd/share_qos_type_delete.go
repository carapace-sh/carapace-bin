package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_qos_type_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a qos type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_qos_type_deleteCmd).Standalone()

	share_qos_typeCmd.AddCommand(share_qos_type_deleteCmd)
}
