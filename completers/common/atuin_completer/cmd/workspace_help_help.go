package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_helpCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_helpCmd)
}
