package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Commands to manage workspace version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_versionCmd).Standalone()

	workspaceCmd.AddCommand(workspace_versionCmd)
}
