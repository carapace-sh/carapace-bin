package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var permissionResetCmd = &cobra.Command{
	Use:     "permission-reset [OPTION…] APP_ID",
	Short:   "Reset permissions for an app",
	GroupID: "permission",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(permissionResetCmd).Standalone()

	permissionResetCmd.Flags().Bool("all", false, "Reset all permissions")
	permissionResetCmd.Flags().BoolP("help", "h", false, "Show help options")
	permissionResetCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	permissionResetCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(permissionResetCmd)

	carapace.Gen(permissionResetCmd).PositionalCompletion(
		flatpak.ActionApplications(),
	)
}
