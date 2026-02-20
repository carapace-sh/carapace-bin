package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_platform_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the platforms in the workspace file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_platform_help_listCmd).Standalone()

	workspace_platform_helpCmd.AddCommand(workspace_platform_help_listCmd)
}
