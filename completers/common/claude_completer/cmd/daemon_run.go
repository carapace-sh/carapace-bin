package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the supervisor in the foreground",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_runCmd).Standalone()

	daemonCmd.AddCommand(daemon_runCmd)

	carapace.Gen(daemon_runCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
