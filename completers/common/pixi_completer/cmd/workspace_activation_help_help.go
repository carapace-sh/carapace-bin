package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_activation_help_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print help for subcommands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_activation_help_helpCmd).Standalone()
	workspace_activation_helpCmd.AddCommand(workspace_activation_help_helpCmd)
}
