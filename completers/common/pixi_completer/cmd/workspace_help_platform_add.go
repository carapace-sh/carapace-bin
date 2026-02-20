package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_platform_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a platform(s) to the workspace file and updates the lockfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_platform_addCmd).Standalone()

	workspace_help_platformCmd.AddCommand(workspace_help_platform_addCmd)
}
