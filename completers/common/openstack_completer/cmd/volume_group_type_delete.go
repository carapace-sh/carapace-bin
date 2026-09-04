package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_group_type_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a volume group type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_type_deleteCmd).Standalone()

	volume_group_typeCmd.AddCommand(volume_group_type_deleteCmd)
}
