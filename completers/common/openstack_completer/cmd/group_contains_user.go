package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var group_contains_userCmd = &cobra.Command{
	Use:   "user",
	Short: "Check user membership in group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(group_contains_userCmd).Standalone()

	group_contains_userCmd.Flags().String("group-domain", "", "Domain the group belongs to (name or ID).")
	group_contains_userCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	group_containsCmd.AddCommand(group_contains_userCmd)
}
