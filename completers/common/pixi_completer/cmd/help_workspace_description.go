package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_descriptionCmd = &cobra.Command{
	Use:   "description",
	Short: "Commands to manage workspace description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_descriptionCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_descriptionCmd)
}
