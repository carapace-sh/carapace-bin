package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Commands to manage workspace version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_versionCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_versionCmd)
}
