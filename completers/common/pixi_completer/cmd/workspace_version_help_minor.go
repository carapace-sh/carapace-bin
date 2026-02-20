package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_version_help_minorCmd = &cobra.Command{
	Use:   "minor",
	Short: "Bump the workspace version to MINOR",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_version_help_minorCmd).Standalone()

	workspace_version_helpCmd.AddCommand(workspace_version_help_minorCmd)
}
