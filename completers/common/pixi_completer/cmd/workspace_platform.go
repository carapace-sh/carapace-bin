package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_platformCmd = &cobra.Command{
	Use:   "platform",
	Short: "Commands to manage workspace platforms",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_platformCmd).Standalone()

	workspaceCmd.AddCommand(workspace_platformCmd)
}
