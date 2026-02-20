package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_description_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the workspace description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_description_setCmd).Standalone()

	help_workspace_descriptionCmd.AddCommand(help_workspace_description_setCmd)
}
