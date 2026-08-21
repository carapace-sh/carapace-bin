package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_type_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete volume type(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_type_deleteCmd).Standalone()

	volume_typeCmd.AddCommand(volume_type_deleteCmd)
}
