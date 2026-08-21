package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set security group properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_setCmd).Standalone()

	security_group_setCmd.Flags().String("description", "", "New security group description")
	security_group_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	security_group_setCmd.Flags().String("name", "", "New security group name")
	security_group_setCmd.Flags().Bool("no-tag", false, "Clear tags associated with the security group.")
	security_group_setCmd.Flags().Bool("stateful", false, "Security group is stateful (default)")
	security_group_setCmd.Flags().Bool("stateless", false, "Security group is stateless")
	security_group_setCmd.Flags().String("tag", "", "Tag to be added to the security group (repeat option to set multiple tags)")
	security_groupCmd.AddCommand(security_group_setCmd)
}
