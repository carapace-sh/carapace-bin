package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_dependenciesCmd = &cobra.Command{
	Use:   "dependencies",
	Short: "Commands to manage workspace dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_dependenciesCmd).Standalone()
	workspace_helpCmd.AddCommand(workspace_help_dependenciesCmd)
}
