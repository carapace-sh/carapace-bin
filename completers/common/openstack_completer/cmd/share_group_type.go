package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_typeCmd = &cobra.Command{
	Use:   "type",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_typeCmd).Standalone()

	share_groupCmd.AddCommand(share_group_typeCmd)
}
