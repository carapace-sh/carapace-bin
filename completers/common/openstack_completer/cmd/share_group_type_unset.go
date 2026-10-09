package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_type_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset share group type extra specs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_type_unsetCmd).Standalone()

	share_group_typeCmd.AddCommand(share_group_type_unsetCmd)
}
