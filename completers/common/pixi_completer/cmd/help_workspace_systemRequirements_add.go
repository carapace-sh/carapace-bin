package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_systemRequirements_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds an environment to the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_systemRequirements_addCmd).Standalone()

	help_workspace_systemRequirementsCmd.AddCommand(help_workspace_systemRequirements_addCmd)
}
