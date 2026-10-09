package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_type_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set share group type properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_type_setCmd).Standalone()

	share_group_type_setCmd.Flags().String("group-specs", "", "Extra specs key and value of share group type that will be used for share type creation.")
	share_group_typeCmd.AddCommand(share_group_type_setCmd)
}
