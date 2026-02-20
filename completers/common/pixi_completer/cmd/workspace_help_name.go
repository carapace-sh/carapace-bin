package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Commands to manage workspace name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_nameCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_nameCmd)
}
