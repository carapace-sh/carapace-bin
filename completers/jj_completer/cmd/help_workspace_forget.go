package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Stop tracking a workspace's working-copy commit in the repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_forgetCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_forgetCmd)
}
