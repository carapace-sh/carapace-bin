package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_description_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the workspace description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_description_getCmd).Standalone()

	help_workspace_descriptionCmd.AddCommand(help_workspace_description_getCmd)
}
