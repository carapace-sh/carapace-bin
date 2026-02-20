package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_description_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the workspace description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_description_getCmd).Standalone()

	workspace_descriptionCmd.AddCommand(workspace_description_getCmd)
}
