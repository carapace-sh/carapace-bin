package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_version_patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Bump the workspace version to PATCH",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_version_patchCmd).Standalone()

	workspace_help_versionCmd.AddCommand(workspace_help_version_patchCmd)
}
