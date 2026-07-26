package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_platform_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List the platforms in the workspace file",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_platform_listCmd).Standalone()

	workspace_platform_listCmd.Flags().Bool("json", false, "Emit machine-readable JSON instead of the human view")
	workspace_platformCmd.AddCommand(workspace_platform_listCmd)
}
