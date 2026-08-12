package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_dependency_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add dependency edges from a source file to one or more dependency files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_dependency_addCmd).Standalone()

	file_help_dependencyCmd.AddCommand(file_help_dependency_addCmd)
}
