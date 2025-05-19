package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:     "prune",
	Short:   "Removes extraneous packages",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()

	pruneCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	pruneCmd.Flags().Bool("color", false, "Controls colors in the output")
	pruneCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	pruneCmd.Flags().BoolP("help", "h", false, "Output usage information")
	pruneCmd.Flags().String("loglevel", "", "What level of logs to report")
	pruneCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	pruneCmd.Flags().Bool("no-optional", false, "Remove the packages specified in `optionalDependencies`")
	pruneCmd.Flags().Bool("prod", false, "Remove the packages specified in `devDependencies`")
	pruneCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	pruneCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	pruneCmd.Flags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.AddCommand(pruneCmd)

	carapace.Gen(pruneCmd).FlagCompletion(carapace.ActionMap{
		"dir":      carapace.ActionDirectories(),
		"loglevel": pnpm.ActionLoglevel(),
	})
}
