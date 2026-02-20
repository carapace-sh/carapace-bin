package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_environment_removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove an environment from the manifest file",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_environment_removeCmd).Standalone()

	workspace_environmentCmd.AddCommand(workspace_environment_removeCmd)
}
