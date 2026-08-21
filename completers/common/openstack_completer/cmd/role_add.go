package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var role_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a role assignment to a user or group on the system, a domain, or a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(role_addCmd).Standalone()

	role_addCmd.Flags().String("domain", "", "Include <domain> (name or ID)")
	role_addCmd.Flags().String("group", "", "Include <group> (name or ID)")
	role_addCmd.Flags().String("group-domain", "", "Domain the group belongs to (name or ID).")
	role_addCmd.Flags().Bool("inherited", false, "Specifies if the role grant is inheritable to the sub projects")
	role_addCmd.Flags().String("project", "", "Include <project> (name or ID)")
	role_addCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	role_addCmd.Flags().String("role-domain", "", "Domain the role belongs to (name or ID).")
	role_addCmd.Flags().String("system", "", "Include <system> (all)")
	role_addCmd.Flags().String("user", "", "Include <user> (name or ID)")
	role_addCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	roleCmd.AddCommand(role_addCmd)
}
