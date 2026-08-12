package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_dependency_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add dependency edges from a source file to one or more dependency files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_dependency_addCmd).Standalone()

	help_file_dependencyCmd.AddCommand(help_file_dependency_addCmd)
}
