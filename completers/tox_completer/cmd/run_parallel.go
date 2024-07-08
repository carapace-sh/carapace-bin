package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runParallelCmd = &cobra.Command{
	Use:   "run-parallel",
	Aliases: []string{"p"},
	Short: "run environments in parallel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runParallelCmd).Standalone()

	add_common_flags(runParallelCmd)
	add_common_subcommand_flags(runParallelCmd)
	add_pkg_only_flags(runParallelCmd)
	add_env_filtering_flags(runParallelCmd)
	add_env_select_flag(runParallelCmd)
	add_common_run_flags(runParallelCmd)

	runParallelCmd.Flags().StringP("parallel", "p", "auto", "run tox environments in parallel, the argument controls limit: all, auto - cpu count, some positive number, zero is turn off (default: auto)")
	runParallelCmd.Flags().BoolP("parallel-live", "o", false, "connect to stdout while running environments")
	runParallelCmd.Flags().Bool("parallel-no-spinner", false, "run tox environments in parallel, but don't show the spinner, implies --parallel")

	carapace.Gen(runParallelCmd).FlagCompletion(carapace.ActionMap{
		"parallel":  carapace.ActionValues("all", "auto", "0"),  // or any positive integer
	})

	rootCmd.AddCommand(runParallelCmd)
}
