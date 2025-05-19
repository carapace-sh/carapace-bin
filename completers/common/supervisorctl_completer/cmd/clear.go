package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/supervisor"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear a process’ log files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clearCmd).Standalone()

	rootCmd.AddCommand(clearCmd)

	carapace.Gen(clearCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				supervisor.ActionProcesses(rootCmd.Flag("configuration").Value.String()).FilterArgs(),
				carapace.ActionValues("all"),
			).ToA()
		}),
	)
}
