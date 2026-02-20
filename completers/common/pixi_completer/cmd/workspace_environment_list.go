package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_environment_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List the environments in the manifest file",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_environment_listCmd).Standalone()

	workspace_environmentCmd.AddCommand(workspace_environment_listCmd)
}
