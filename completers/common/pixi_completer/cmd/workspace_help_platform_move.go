package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_platform_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Reorder a workspace platform, changing its selection priority",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_platform_moveCmd).Standalone()

	workspace_help_platformCmd.AddCommand(workspace_help_platform_moveCmd)
}
