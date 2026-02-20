package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_name_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the workspace name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_name_setCmd).Standalone()

	workspace_help_nameCmd.AddCommand(workspace_help_name_setCmd)
}
