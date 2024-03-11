package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var permissionShowCmd = &cobra.Command{
	Use:     "permission-show [OPTION…] APP_ID",
	Short:   "Show permissions for an app",
	GroupID: "permission",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(permissionShowCmd).Standalone()

	permissionShowCmd.Flags().BoolP("help", "h", false, "Show help options")
	permissionShowCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	permissionShowCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(permissionShowCmd)

	carapace.Gen(permissionShowCmd).PositionalCompletion(
		flatpak.ActionApplications(),
	)
}
