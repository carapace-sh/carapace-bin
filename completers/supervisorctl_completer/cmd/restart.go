package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/supervisor"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restartCmd).Standalone()

	rootCmd.AddCommand(restartCmd)

	carapace.Gen(restartCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				supervisor.ActionGroups(rootCmd.Flag("configuration").Value.String()),
				supervisor.ActionProcesses(rootCmd.Flag("configuration").Value.String()),
				carapace.ActionValues("all"),
			).ToA().FilterArgs()
		}),
	)
}
