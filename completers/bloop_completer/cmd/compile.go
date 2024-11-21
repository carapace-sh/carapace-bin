package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "compile projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compileCmd).Standalone()

	compileCmd.Flags().Bool("cascade", false, "compile a project and all projects depending on it")
	compileCmd.Flags().Bool("incremental", false, "compile the project incrementally")
	compileCmd.Flags().Bool("pipeline", false, "pipeline the compilation of modules in your build")
	compileCmd.Flags().StringSliceP("project", "p", nil, "the projects to compile")
	compileCmd.Flags().StringSlice("projects", nil, "the projects to compile")
	compileCmd.Flags().String("reporter", "", "pick reporter to show compilation messages")
	compileCmd.Flags().BoolP("watch", "w", false, "run the command when projects' source files change")
	rootCmd.AddCommand(compileCmd)

	// TODO completion
}
