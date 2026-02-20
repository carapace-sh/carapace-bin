package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_systemRequirements_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the environments in the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_systemRequirements_help_listCmd).Standalone()

	workspace_systemRequirements_helpCmd.AddCommand(workspace_systemRequirements_help_listCmd)
}
