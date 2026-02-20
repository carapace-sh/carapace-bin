package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_environment_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the environments in the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_environment_listCmd).Standalone()

	workspace_help_environmentCmd.AddCommand(workspace_help_environment_listCmd)
}
