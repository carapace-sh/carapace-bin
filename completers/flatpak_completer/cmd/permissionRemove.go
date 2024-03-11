package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var permissionRemoveCmd = &cobra.Command{
	Use:     "permission-remove [OPTION…] TABLE ID [APP_ID]",
	Short:   "Remove item from permission store",
	GroupID: "permission",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(permissionRemoveCmd).Standalone()

	permissionRemoveCmd.Flags().BoolP("help", "h", false, "Show help options")
	permissionRemoveCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	permissionRemoveCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(permissionRemoveCmd)

	carapace.Gen(permissionRemoveCmd).PositionalCompletion(
		carapace.ActionValues(), // TODO table
		carapace.ActionValues(), // TODO id
		flatpak.ActionApplications(),
	)
}
