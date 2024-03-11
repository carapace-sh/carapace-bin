package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/supervisor"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stopCmd).Standalone()

	rootCmd.AddCommand(stopCmd)

	carapace.Gen(stopCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				supervisor.ActionGroups(rootCmd.Flag("configuration").Value.String()),
				supervisor.ActionProcesses(rootCmd.Flag("configuration").Value.String()),
				carapace.ActionValues("all"),
			).ToA().FilterArgs()
		}),
	)
}
