package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_type_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set share type properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_type_setCmd).Standalone()

	share_type_setCmd.Flags().String("description", "", "New description of share type.")
	share_type_setCmd.Flags().String("extra-specs", "", "Extra specs key and value of share type that will be used for share type creation.")
	share_type_setCmd.Flags().String("name", "", "New name of share type.")
	share_type_setCmd.Flags().String("public", "", "New visibility of the share type.")
	share_typeCmd.AddCommand(share_type_setCmd)
}
