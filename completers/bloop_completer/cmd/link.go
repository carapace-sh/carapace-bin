package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bloop"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:   "link <link>",
	Short: "link projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	linkCmd.Flags().Bool("incremental", false, "compile the project incrementally")
	linkCmd.Flags().StringP("main", "m", "", "the main class to link")
	linkCmd.Flags().StringP("optimize", "O", "", "optimization level of the linker")
	linkCmd.Flags().Bool("pipeline", false, "pipeline the compilation of modules in your build")
	linkCmd.Flags().StringSliceP("project", "p", nil, "the projects to link")
	linkCmd.Flags().StringSlice("projects", nil, "the projects to link")
	linkCmd.Flags().Bool("reporter", false, "pick reporter to show compilation messages")
	linkCmd.Flags().BoolP("watch", "w", false, "run the command whenever projects' source files change")
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).FlagCompletion(carapace.ActionMap{
		"main":     carapace.ActionValues(),
		"optimize": carapace.ActionValues("debug", "release"),
		"project":  bloop.ActionProjects(),
		"projects": bloop.ActionProjects(),
	})

	carapace.Gen(linkCmd).PositionalAnyCompletion(
		bloop.ActionProjects().FilterArgs(),
	)
}
