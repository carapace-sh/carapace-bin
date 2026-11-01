package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_type_accessCmd = &cobra.Command{
	Use:   "access",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_type_accessCmd).Standalone()

	share_group_typeCmd.AddCommand(share_group_type_accessCmd)
}
