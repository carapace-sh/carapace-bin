package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var user_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete user(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(user_deleteCmd).Standalone()

	user_deleteCmd.Flags().String("domain", "", "Domain owning <user> (name or ID)")
	userCmd.AddCommand(user_deleteCmd)
}
