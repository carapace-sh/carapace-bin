package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_helpCmd).Standalone()

	workspaceCmd.AddCommand(workspace_helpCmd)
}
