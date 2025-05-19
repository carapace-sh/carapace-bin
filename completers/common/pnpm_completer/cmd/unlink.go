package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:     "unlink",
	Aliases: []string{"dislink"},
	Short:   "Removes the link created by `pnpm link` and reinstalls package if it is saved in `package.json`",
	GroupID: "manage",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlinkCmd).Standalone()

	unlinkCmd.Flags().Bool("aggregate-output", false, "Aggregate output from child processes that are run in parallel")
	unlinkCmd.Flags().Bool("color", false, "Controls colors in the output")
	unlinkCmd.Flags().StringP("dir", "C", "", "Change to directory <dir>")
	unlinkCmd.Flags().BoolP("help", "h", false, "Output usage information")
	unlinkCmd.Flags().String("loglevel", "", "What level of logs to report")
	unlinkCmd.Flags().Bool("no-color", false, "Controls colors in the output")
	unlinkCmd.Flags().BoolP("recursive", "r", false, "Unlink in every package found in subdirectories or in every workspace package")
	unlinkCmd.Flags().Bool("stream", false, "Stream output from child processes immediately")
	unlinkCmd.Flags().Bool("use-stderr", false, "Divert all output to stderr")
	unlinkCmd.Flags().BoolP("workspace-root", "w", false, "Run the command on the root workspace project")
	rootCmd.AddCommand(unlinkCmd)

	carapace.Gen(unlinkCmd).FlagCompletion(carapace.ActionMap{
		"dir":      carapace.ActionDirectories(),
		"loglevel": pnpm.ActionLoglevel(),
	})

	carapace.Gen(unlinkCmd).PositionalAnyCompletion(
		pnpm.ActionDependencyNames(), // TODO complete only linked dependencies
	)
}
