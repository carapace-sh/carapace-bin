package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_version_majorCmd = &cobra.Command{
	Use:   "major",
	Short: "Bump the workspace version to MAJOR",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_version_majorCmd).Standalone()

	workspace_help_versionCmd.AddCommand(workspace_help_version_majorCmd)
}
