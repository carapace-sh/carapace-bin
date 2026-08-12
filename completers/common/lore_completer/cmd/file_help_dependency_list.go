package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_dependency_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List dependencies or dependents for files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_dependency_listCmd).Standalone()

	file_help_dependencyCmd.AddCommand(file_help_dependency_listCmd)
}
