package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Show the current workspace root directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_rootCmd).Standalone()

	workspace_helpCmd.AddCommand(workspace_help_rootCmd)
}
