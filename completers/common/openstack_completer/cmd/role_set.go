package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var role_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set role properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(role_setCmd).Standalone()

	role_setCmd.Flags().String("description", "", "Add description about the role")
	role_setCmd.Flags().String("domain", "", "Domain the role belongs to (name or ID)")
	role_setCmd.Flags().Bool("immutable", false, "Make resource immutable.")
	role_setCmd.Flags().String("name", "", "Set role name")
	role_setCmd.Flags().Bool("no-immutable", false, "Make resource mutable (default)")
	roleCmd.AddCommand(role_setCmd)
}
