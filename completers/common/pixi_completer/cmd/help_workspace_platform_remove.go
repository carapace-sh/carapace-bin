package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_platform_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove platform(s) from the workspace file and updates the lockfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_platform_removeCmd).Standalone()

	help_workspace_platformCmd.AddCommand(help_workspace_platform_removeCmd)
}
