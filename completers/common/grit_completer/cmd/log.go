package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show recent commits reachable from HEAD",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()

	logCmd.Flags().String("before", "", "Continue listing from before this commit (for paging)")
	logCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(logCmd)

	carapace.Gen(logCmd).FlagCompletion(carapace.ActionMap{
		"before": git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(logCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			git.ActionRefRanges(git.RefOption{}.Default()).UnlessF(condition.CompletingPath),
		).ToA(),
	)

	carapace.Gen(logCmd).DashAnyCompletion(
		carapace.ActionFiles(),
	)
}
