package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset security group properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_unsetCmd).Standalone()

	security_group_unsetCmd.Flags().Bool("all-tag", false, "Clear all tags associated with the security group")
	security_group_unsetCmd.Flags().String("tag", "", "Tag to be removed from the security group (repeat option to remove multiple tags)")
	security_groupCmd.AddCommand(security_group_unsetCmd)
}
