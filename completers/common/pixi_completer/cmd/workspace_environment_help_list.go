package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_environment_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the environments in the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_environment_help_listCmd).Standalone()

	workspace_environment_helpCmd.AddCommand(workspace_environment_help_listCmd)
}
