package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/devenv_completer/cmd/action"
	"github.com/spf13/cobra"
)

var processes_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a process, honoring its dependencies (or all processes if no name given)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processes_startCmd).Standalone()

	processes_startCmd.Flags().BoolP("detach", "d", false, "Start processes in the background")

	processesCmd.AddCommand(processes_startCmd)

	carapace.Gen(processes_startCmd).PositionalCompletion(
		action.ActionProcesses(processes_startCmd),
	)
}
