package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_version_minorCmd = &cobra.Command{
	Use:   "minor",
	Short: "Bump the workspace version to MINOR",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_version_minorCmd).Standalone()

	workspace_versionCmd.AddCommand(workspace_version_minorCmd)
}
