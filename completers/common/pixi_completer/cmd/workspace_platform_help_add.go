package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_platform_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a platform(s) to the workspace file and updates the lockfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_platform_help_addCmd).Standalone()

	workspace_platform_helpCmd.AddCommand(workspace_platform_help_addCmd)
}
