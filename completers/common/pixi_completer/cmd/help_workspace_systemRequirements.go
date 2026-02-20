package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_systemRequirementsCmd = &cobra.Command{
	Use:   "system-requirements",
	Short: "Commands to manage workspace system requirements",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_systemRequirementsCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_systemRequirementsCmd)
}
