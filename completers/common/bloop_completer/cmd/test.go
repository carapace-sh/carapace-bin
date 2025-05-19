package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bloop"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test <project>",
	Short: "run tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().StringSlice("args", nil, "the arguments to pass in to the test framework")
	testCmd.Flags().Bool("cascade", false, "test a project and all projects depending on it")
	testCmd.Flags().Bool("include-dependencies", false, "test a project and all its dependencies")
	testCmd.Flags().Bool("incremental", false, "compile the project incrementally")
	testCmd.Flags().StringP("only", "o", "", "the list of test suite filters to test for only")
	testCmd.Flags().Bool("parallel", false, "run tests in parallel")
	testCmd.Flags().Bool("pipeline", false, "pipeline the compilation of modules in your build")
	testCmd.Flags().StringSliceP("project", "p", nil, "the projects to test")
	testCmd.Flags().StringSlice("projects", nil, "the projects to test")
	testCmd.Flags().Bool("propagate", false, "test a project and all its dependencies")
	testCmd.Flags().String("reporter", "", "pick reporter to show compilation messages")
	testCmd.Flags().BoolP("watch", "w", false, "run the command when projects' source files change")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"args":     carapace.ActionValues(),
		"only":     carapace.ActionValues(),
		"project":  bloop.ActionProjects(),
		"projects": bloop.ActionProjects(),
		"reporter": bloop.ActionReporters(),
	})

	carapace.Gen(testCmd).PositionalAnyCompletion(
		bloop.ActionProjects().FilterArgs(),
	)
}
