package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_dependency_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add dependency edges from a source file to one or more dependency files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_dependency_addCmd).Standalone()

	file_dependency_addCmd.Flags().Bool("force", false, "Skip cycle detection")
	file_dependency_addCmd.Flags().BoolP("help", "h", false, "Print help")
	file_dependency_addCmd.Flags().StringSlice("tag", nil, "Tags to apply to all added dependency edges")
	file_dependencyCmd.AddCommand(file_dependency_addCmd)
}
