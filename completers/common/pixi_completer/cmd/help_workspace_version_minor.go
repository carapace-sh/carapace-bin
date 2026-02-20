package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_version_minorCmd = &cobra.Command{
	Use:   "minor",
	Short: "Bump the workspace version to MINOR",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_version_minorCmd).Standalone()

	help_workspace_versionCmd.AddCommand(help_workspace_version_minorCmd)
}
