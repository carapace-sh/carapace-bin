package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_version_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the workspace version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_version_getCmd).Standalone()

	workspace_help_versionCmd.AddCommand(workspace_help_version_getCmd)
}
