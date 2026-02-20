package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_descriptionCmd = &cobra.Command{
	Use:   "description",
	Short: "Commands to manage workspace description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_descriptionCmd).Standalone()

	workspaceCmd.AddCommand(workspace_descriptionCmd)
}
