package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_version_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the workspace version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_version_help_setCmd).Standalone()

	workspace_version_helpCmd.AddCommand(workspace_version_help_setCmd)
}
