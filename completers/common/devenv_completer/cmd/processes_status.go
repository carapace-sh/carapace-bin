package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/devenv_completer/cmd/action"
	"github.com/spf13/cobra"
)

var processes_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status of a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processes_statusCmd).Standalone()

	processesCmd.AddCommand(processes_statusCmd)

	carapace.Gen(processes_statusCmd).PositionalCompletion(
		action.ActionRunningProcesses(processes_statusCmd),
	)
}
