package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Show the current workspace root directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_rootCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_rootCmd)
}
