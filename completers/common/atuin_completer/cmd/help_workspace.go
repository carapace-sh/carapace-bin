package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Set up the .atuin.lua config file of a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspaceCmd).Standalone()

	helpCmd.AddCommand(help_workspaceCmd)
}
