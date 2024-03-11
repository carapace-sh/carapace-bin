package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var permissionSetCmd = &cobra.Command{
	Use:     "permission-set [OPTION…] TABLE ID APP_ID [PERMISSION...]",
	Short:   "Set permissions",
	GroupID: "permission",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(permissionSetCmd).Standalone()

	permissionSetCmd.Flags().String("data", "", "Associate DATA with the entry")
	permissionSetCmd.Flags().BoolP("help", "h", false, "Show help options")
	permissionSetCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	permissionSetCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(permissionSetCmd)

	carapace.Gen(permissionSetCmd).PositionalCompletion(
		carapace.ActionValues(), // TODO table
		carapace.ActionValues(), // TODO id
		flatpak.ActionApplications(),
	)

	// TODO permissions
}
