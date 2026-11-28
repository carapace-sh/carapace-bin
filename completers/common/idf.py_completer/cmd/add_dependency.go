package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addDependencyCmd = &cobra.Command{
	Use:   "add-dependency",
	Short: "Add dependency to the manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addDependencyCmd).Standalone()

	addDependencyCmd.Flags().String("component", "", "Name of the component in the project")
	addDependencyCmd.Flags().String("git", "", "Git URL of the component")
	addDependencyCmd.Flags().String("git-path", "", "Path to the component in the git repository")
	addDependencyCmd.Flags().String("git-ref", "", "Git reference (branch, tag, commit SHA)")
	addDependencyCmd.Flags().String("path", "", "Path to the component where the dependency will be added")
	rootCmd.AddCommand(addDependencyCmd)

	carapace.Gen(addDependencyCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})
}
