package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/supervisor"
	"github.com/spf13/cobra"
)

var pidCmd = &cobra.Command{
	Use:   "pid",
	Short: "Get the PID of supervisord",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pidCmd).Standalone()

	rootCmd.AddCommand(pidCmd)

	carapace.Gen(pidCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				supervisor.ActionProcesses(rootCmd.Flag("configuration").Value.String()),
				carapace.ActionValues("all"),
			).ToA()
		}),
	)
}
