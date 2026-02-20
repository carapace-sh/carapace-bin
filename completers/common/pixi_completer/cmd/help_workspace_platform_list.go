package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_platform_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the platforms in the workspace file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_platform_listCmd).Standalone()

	help_workspace_platformCmd.AddCommand(help_workspace_platform_listCmd)
}
