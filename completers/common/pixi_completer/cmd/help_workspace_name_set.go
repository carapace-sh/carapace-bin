package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_name_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the workspace name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_name_setCmd).Standalone()

	help_workspace_nameCmd.AddCommand(help_workspace_name_setCmd)
}
