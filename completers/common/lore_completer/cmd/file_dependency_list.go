package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_dependency_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List dependencies or dependents for files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_dependency_listCmd).Standalone()

	file_dependency_listCmd.Flags().String("depth", "0", "Maximum recursion depth (0 = unlimited)")
	file_dependency_listCmd.Flags().BoolP("help", "h", false, "Print help")
	file_dependency_listCmd.Flags().Bool("recursive", false, "Recursively resolve transitive dependencies")
	file_dependency_listCmd.Flags().Bool("reverse", false, "List dependents instead of dependencies")
	file_dependency_listCmd.Flags().String("revision", "", "Revision to query (defaults to staged/current)")
	file_dependency_listCmd.Flags().StringSlice("tag", nil, "Filter by tag")
	file_dependencyCmd.AddCommand(file_dependency_listCmd)
}
