package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a share group property",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_unsetCmd).Standalone()

	share_group_unsetCmd.Flags().Bool("description", false, "Unset share group description.")
	share_group_unsetCmd.Flags().Bool("name", false, "Unset share group name.")
	share_groupCmd.AddCommand(share_group_unsetCmd)
}
