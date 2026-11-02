package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().StringS("C", "C", "", "directory to cd into before running")
	initCmd.Flags().BoolP("build", "b", false, "build after generation")
	initCmd.Flags().String("builddir", "", "directory for build")
	initCmd.Flags().StringP("deps", "d", "", "dependencies, comma-separated")
	initCmd.Flags().StringP("executable", "e", "", "executable name.")
	initCmd.Flags().BoolP("force", "f", false, "force overwrite of existing files and directories.")
	initCmd.Flags().StringP("language", "l", "", "project language.")
	initCmd.Flags().StringP("name", "n", "", "project name.")
	initCmd.Flags().String("type", "", "project type.")
	initCmd.Flags().String("version", "", "project version.")
	applyCommonFlagCompletion(initCmd)
	rootCmd.AddCommand(initCmd)
}
