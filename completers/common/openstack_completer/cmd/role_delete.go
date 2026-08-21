package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var role_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete role(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(role_deleteCmd).Standalone()

	role_deleteCmd.Flags().String("domain", "", "Domain the role belongs to (name or ID)")
	roleCmd.AddCommand(role_deleteCmd)
}
