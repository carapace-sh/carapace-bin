package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var role_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a role assignment from system/domain/project : user/group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(role_removeCmd).Standalone()

	role_removeCmd.Flags().String("domain", "", "Include <domain> (name or ID)")
	role_removeCmd.Flags().String("group", "", "Include <group> (name or ID)")
	role_removeCmd.Flags().String("group-domain", "", "Domain the group belongs to (name or ID).")
	role_removeCmd.Flags().Bool("inherited", false, "Specifies if the role grant is inheritable to the sub projects")
	role_removeCmd.Flags().String("project", "", "Include <project> (name or ID)")
	role_removeCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	role_removeCmd.Flags().String("role-domain", "", "Domain the role belongs to (name or ID).")
	role_removeCmd.Flags().String("system", "", "Include <system> (all)")
	role_removeCmd.Flags().String("user", "", "Include <user> (name or ID)")
	role_removeCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	roleCmd.AddCommand(role_removeCmd)
}
