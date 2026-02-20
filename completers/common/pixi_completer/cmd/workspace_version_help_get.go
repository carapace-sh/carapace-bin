package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_version_help_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the workspace version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_version_help_getCmd).Standalone()

	workspace_version_helpCmd.AddCommand(workspace_version_help_getCmd)
}
