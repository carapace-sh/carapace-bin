package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listPermissionGroupsCmd = &cobra.Command{
	Use:   "permission-groups",
	Short: "Print all known permission groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPermissionGroupsCmd).Standalone()

	listCmd.AddCommand(listPermissionGroupsCmd)

}
