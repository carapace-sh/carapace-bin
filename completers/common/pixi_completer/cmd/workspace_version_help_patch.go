package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_version_help_patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Bump the workspace version to PATCH",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_version_help_patchCmd).Standalone()

	workspace_version_helpCmd.AddCommand(workspace_version_help_patchCmd)
}
