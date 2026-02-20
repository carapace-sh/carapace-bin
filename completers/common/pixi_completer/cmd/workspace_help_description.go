package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_descriptionCmd = &cobra.Command{
	Use:   "description",
	Short: "Commands to manage workspace description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_descriptionCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_descriptionCmd)
}
