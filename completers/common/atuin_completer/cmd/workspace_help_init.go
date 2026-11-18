package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a .atuin.lua workspace config file in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_initCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_initCmd)
}
