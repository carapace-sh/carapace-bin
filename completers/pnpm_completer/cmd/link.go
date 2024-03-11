package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:     "link",
	Aliases: []string{"ln"},
	Short:   "Connect the local project to another one",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	linkCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	linkCmd.Flags().Bool("color", false, "Controls colors in the output")
	linkCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	linkCmd.Flags().BoolP("help", "h", false, "Output usage information")
	linkCmd.Flags().String("loglevel", "", "What level of logs to report")
	linkCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	linkCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	linkCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	linkCmd.Flags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).FlagCompletion(carapace.ActionMap{
		"dir":      carapace.ActionDirectories(),
		"loglevel": pnpm.ActionLoglevel(),
	})

	carapace.Gen(linkCmd).PositionalCompletion(
		carapace.Batch(
			pnpm.ActionDependencyNames(),
			carapace.ActionDirectories(),
		).ToA(),
	)
}
