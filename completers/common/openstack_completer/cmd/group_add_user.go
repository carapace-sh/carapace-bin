package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_add_userCmd = &cobra.Command{
	Use:   "user",
	Short: "Add user to group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_add_userCmd).Standalone()

	group_add_userCmd.Flags().String("group-domain", "", "Domain the group belongs to (name or ID).")
	group_add_userCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	group_addCmd.AddCommand(group_add_userCmd)
}
