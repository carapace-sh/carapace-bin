package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_dependencies_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print help for subcommands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_dependencies_helpCmd).Standalone()
	workspace_dependenciesCmd.AddCommand(workspace_dependencies_helpCmd)
}
