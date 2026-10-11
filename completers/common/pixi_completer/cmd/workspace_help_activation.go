package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_activationCmd = &cobra.Command{
	Use:   "activation",
	Short: "Commands to manage environment activation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_activationCmd).Standalone()
	workspace_helpCmd.AddCommand(workspace_help_activationCmd)
}
