package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_name_help_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the workspace name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_name_help_setCmd).Standalone()

	workspace_name_helpCmd.AddCommand(workspace_name_help_setCmd)
}
