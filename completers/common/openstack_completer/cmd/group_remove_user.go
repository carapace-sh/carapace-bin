package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_remove_userCmd = &cobra.Command{
	Use:   "user",
	Short: "Remove user from group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_remove_userCmd).Standalone()

	group_remove_userCmd.Flags().String("group-domain", "", "Domain the group belongs to (name or ID).")
	group_remove_userCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	group_removeCmd.AddCommand(group_remove_userCmd)
}
