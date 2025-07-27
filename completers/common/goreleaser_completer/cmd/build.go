package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/goreleaser"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build",
	Short:   "Builds the current project",
	Aliases: []string{"b"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("auto-snapshot", false, "Automatically sets --snapshot if the repository is dirty")
	buildCmd.Flags().Bool("clean", false, "Removes the 'dist' directory before building")
	buildCmd.Flags().StringP("config", "f", "", "Load configuration from file")
	buildCmd.Flags().Bool("deprecated", false, "Force print the deprecation message - tests only")
	buildCmd.Flags().StringSlice("id", nil, "Builds only the specified build ids")
	buildCmd.Flags().StringP("output", "o", "", "Copy the binary to the path after the build. Only taken into account when using --single-target and a single id (either with --id or if configuration only has one build)")
	buildCmd.Flags().StringP("parallelism", "p", "", "Number of tasks to run concurrently (default: number of CPUs)")
	buildCmd.Flags().Bool("single-target", false, "Builds only for current GOOS and GOARCH, regardless of what's set in the configuration file")
	buildCmd.Flags().StringSlice("skip", nil, "Skip the given options (valid options are: before, post-hooks, pre-hooks, validate)")
	buildCmd.Flags().Bool("snapshot", false, "Generate an unversioned snapshot build, skipping all validations")
	buildCmd.Flags().String("timeout", "", "Timeout to the entire build process")
	buildCmd.Flag("deprecated").Hidden = true
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"id": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return goreleaser.ActionBuilds(buildCmd.Flag("config").Value.String()).UniqueList(",")
		}),
		"output": carapace.ActionFiles(),
		"skip":   carapace.ActionValues("before", "post-hooks", "pre-hooks", "validate").UniqueList(","),
	})
}
