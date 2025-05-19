package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bloop"
	"github.com/spf13/cobra"
)

var consoleCmd = &cobra.Command{
	Use:   "console <project>",
	Short: "start console",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consoleCmd).Standalone()

	consoleCmd.Flags().String("ammonite-version", "", "the Ammonite version to use")
	consoleCmd.Flags().StringSlice("args", nil, "the arguments to pass in to Ammonite")
	consoleCmd.Flags().Bool("exclude-root", false, "start up the console compiling only the target project's dependencies")
	consoleCmd.Flags().Bool("incremental", false, "compile the project incrementally")
	consoleCmd.Flags().String("out-file", "", "the output file where the Ammonite command is written")
	consoleCmd.Flags().Bool("pipeline", false, "pipeline the compilation of modules in your build")
	consoleCmd.Flags().StringSliceP("project", "p", nil, "the projects to run the console at")
	consoleCmd.Flags().StringSlice("projects", nil, "the projects to run the console at")
	consoleCmd.Flags().String("repl", "", "pick REPL to run console")
	consoleCmd.Flags().String("reporter", "", "pick reporter to show compilation messages")
	rootCmd.AddCommand(consoleCmd)

	carapace.Gen(consoleCmd).FlagCompletion(carapace.ActionMap{
		"ammonite-version": carapace.ActionValues(),
		"args":             carapace.ActionValues(),
		"out-file":         carapace.ActionFiles(),
		"project":          bloop.ActionProjects(),
		"projects":         bloop.ActionProjects(),
		"repl":             carapace.ActionValues("scalac", "ammonite"),
		"reporter":         bloop.ActionReporters(),
	})

	carapace.Gen(consoleCmd).PositionalCompletion(
		bloop.ActionProjects(),
	)
}
