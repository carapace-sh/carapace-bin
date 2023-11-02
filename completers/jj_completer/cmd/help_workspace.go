package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Commands for working with workspaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspaceCmd).Standalone()

	helpCmd.AddCommand(help_workspaceCmd)
}
