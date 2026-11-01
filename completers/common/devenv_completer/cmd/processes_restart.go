package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/devenv_completer/cmd/action"
	"github.com/spf13/cobra"
)

var processes_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processes_restartCmd).Standalone()

	processesCmd.AddCommand(processes_restartCmd)

	carapace.Gen(processes_restartCmd).PositionalCompletion(
		action.ActionRunningProcesses(processes_restartCmd),
	)
}
