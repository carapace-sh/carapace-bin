package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_version_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the workspace version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_version_setCmd).Standalone()

	workspace_help_versionCmd.AddCommand(workspace_help_version_setCmd)
}
