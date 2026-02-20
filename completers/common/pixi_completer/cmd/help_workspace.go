package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Modify the workspace configuration file through the command line",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspaceCmd).Standalone()

	helpCmd.AddCommand(help_workspaceCmd)
}
