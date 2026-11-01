package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var listPermissionsCmd = &cobra.Command{
	Use:   "permissions",
	Short: "Print all known permissions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPermissionsCmd).Standalone()

	listPermissionsCmd.Flags().BoolP("dangerous", "d", false, "only list dangerous permissions")
	listPermissionsCmd.Flags().BoolP("full", "f", false, "print all information")
	listPermissionsCmd.Flags().BoolP("group", "g", false, "organize by group")
	listPermissionsCmd.Flags().BoolP("short", "s", false, "short summary")
	listPermissionsCmd.Flags().BoolP("user", "u", false, "list only the permissions users will see")

	listCmd.AddCommand(listPermissionsCmd)

	carapace.Gen(listPermissionsCmd).PositionalCompletion(
		android.ActionPermissionGroups(),
	)
}
