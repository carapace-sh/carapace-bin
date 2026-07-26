package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_platform_moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Reorder a workspace platform, changing its selection priority",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_platform_moveCmd).Standalone()

	help_workspace_platformCmd.AddCommand(help_workspace_platform_moveCmd)
}
