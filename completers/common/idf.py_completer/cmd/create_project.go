package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Create a new project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createProjectCmd).Standalone()
	createProjectCmd.Flags().Bool("cpp", false, "Create a C++ main source file instead of C")
	createProjectCmd.Flags().String("path", "", "Set the path for the new project")
	carapace.Gen(createProjectCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionDirectories(),
	})
	rootCmd.AddCommand(createProjectCmd)
}
