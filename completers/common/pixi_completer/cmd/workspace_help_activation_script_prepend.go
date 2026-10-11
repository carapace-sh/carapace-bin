package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_activation_script_prependCmd = &cobra.Command{
	Use:   "prepend",
	Short: "Move activation scripts to the front of the list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_activation_script_prependCmd).Standalone()
	workspace_help_activation_scriptCmd.AddCommand(workspace_help_activation_script_prependCmd)
}
