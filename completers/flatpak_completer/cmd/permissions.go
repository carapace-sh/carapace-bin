package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var permissionsCmd = &cobra.Command{
	Use:     "permissions",
	Short:   "List permissions",
	GroupID: "permission",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(permissionsCmd).Standalone()

	permissionsCmd.Flags().BoolP("help", "h", false, "Show help options")
	permissionsCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	permissionsCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(permissionsCmd)

	// TODO positional
}
