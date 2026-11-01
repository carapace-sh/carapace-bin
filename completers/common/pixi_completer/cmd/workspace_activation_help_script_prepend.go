package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_activation_help_script_prependCmd = &cobra.Command{
	Use:   "prepend",
	Short: "Move activation scripts to the front of the list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_help_script_prependCmd).Standalone()
	workspace_activation_help_scriptCmd.AddCommand(workspace_activation_help_script_prependCmd)
}
