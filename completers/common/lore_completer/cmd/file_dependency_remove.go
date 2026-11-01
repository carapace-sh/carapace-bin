package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_dependency_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove dependency edges from a source file to one or more dependency files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_dependency_removeCmd).Standalone()

	file_dependency_removeCmd.Flags().BoolP("help", "h", false, "Print help")
	file_dependency_removeCmd.Flags().StringSlice("tag", nil, "Remove only specific tags instead of entire edges")
	file_dependencyCmd.AddCommand(file_dependency_removeCmd)

	carapace.Gen(file_dependency_removeCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
