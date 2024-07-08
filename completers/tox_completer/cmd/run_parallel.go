package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var runParallelCmd = &cobra.Command{
	Use:     "run-parallel",
	Aliases: []string{"p"},
	Short:   "run environments in parallel",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runParallelCmd).Standalone()

	runParallelCmd.Flags().StringP("parallel", "p", "auto", "run tox environments in parallel, the argument controls limit: all, auto - cpu count, some positive number, zero is turn off (default: auto)")
	runParallelCmd.Flags().BoolP("parallel-live", "o", false, "connect to stdout while running environments")
	runParallelCmd.Flags().Bool("parallel-no-spinner", false, "run tox environments in parallel, but don't show the spinner, implies --parallel")
	addCommonSubcommandFlags(runParallelCmd)
	addPkgOnlyFlags(runParallelCmd)
	addEnvFilteringFlags(runParallelCmd)
	addEnvSelectFlag(runParallelCmd)
	addCommonRunFlags(runParallelCmd)
	rootCmd.AddCommand(runParallelCmd)

	carapace.Gen(runParallelCmd).FlagCompletion(carapace.ActionMap{
		"parallel": carapace.ActionValues("all", "auto", "0").StyleF(style.ForKeyword), // or any positive integer
	})
}
