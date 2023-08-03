package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/supervisor"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(startCmd).Standalone()

	rootCmd.AddCommand(startCmd)

	carapace.Gen(startCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				supervisor.ActionGroups(rootCmd.Flag("configuration").Value.String()),
				supervisor.ActionProcesses(rootCmd.Flag("configuration").Value.String()),
				carapace.ActionValues("all"),
			).ToA().FilterArgs()
		}),
	)
}
