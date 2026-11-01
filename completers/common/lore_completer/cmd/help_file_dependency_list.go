package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_dependency_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List dependencies or dependents for files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_dependency_listCmd).Standalone()

	help_file_dependencyCmd.AddCommand(help_file_dependency_listCmd)
}
