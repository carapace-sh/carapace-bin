package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/supervisor"
	"github.com/spf13/cobra"
)

var fgCmd = &cobra.Command{
	Use:   "fg",
	Short: "Connect to a process in foreground mode Press Ctrl+C to exit foreground",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fgCmd).Standalone()

	rootCmd.AddCommand(fgCmd)

	carapace.Gen(fgCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return supervisor.ActionProcesses(rootCmd.Flag("configuration").Value.String())
		}),
	)
}
