package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_platformCmd = &cobra.Command{
	Use:   "platform",
	Short: "Commands to manage workspace platforms",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_platformCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_platformCmd)
}
