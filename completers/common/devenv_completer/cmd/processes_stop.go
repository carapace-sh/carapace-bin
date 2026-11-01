package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/devenv_completer/cmd/action"
	"github.com/spf13/cobra"
)

var processes_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a running process (or all processes if no name given)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processes_stopCmd).Standalone()

	processesCmd.AddCommand(processes_stopCmd)

	carapace.Gen(processes_stopCmd).PositionalCompletion(
		action.ActionRunningProcesses(processes_stopCmd),
	)
}
