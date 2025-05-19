package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bloop"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <project>",
	Short: "run a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().StringSlice("args", nil, "the arguments to pass in to the main class")
	runCmd.Flags().Bool("incremental", false, "compile the project incrementally")
	runCmd.Flags().StringP("main", "m", "", "the main class to run")
	runCmd.Flags().StringP("optimize", "O", "", "if an optimizer is used, run it in `debug` or `release` mode")
	runCmd.Flags().Bool("pipeline", false, "pipeline the compilation of modules in your build")
	runCmd.Flags().StringSliceP("project", "p", nil, "the projects to run")
	runCmd.Flags().StringSlice("projects", nil, "the projects to run")
	runCmd.Flags().String("reporter", "", "pick reporter to show compilation messages")
	runCmd.Flags().Bool("skip-jargs", false, "ignore arguments starting with `-J` and forward them instead")
	runCmd.Flags().BoolP("watch", "w", false, "run the command whenever projects' source files change")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"args":     carapace.ActionValues(),
		"main":     carapace.ActionValues(),
		"optimize": carapace.ActionValues("debug", "release"),
		"project":  bloop.ActionProjects(),
		"projects": bloop.ActionProjects(),
		"reporter": bloop.ActionReporters(),
	})

	carapace.Gen(runCmd).PositionalAnyCompletion(
		bloop.ActionProjects().FilterArgs(),
	)
}
