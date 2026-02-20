package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_description_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the workspace description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_description_help_setCmd).Standalone()

	workspace_description_helpCmd.AddCommand(workspace_description_help_setCmd)
}
