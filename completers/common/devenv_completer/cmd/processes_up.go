package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/devenv_completer/cmd/action"
	"github.com/spf13/cobra"
)

var processes_upCmd = &cobra.Command{
	Use:   "up",
	Short: "Start processes in the foreground",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processes_upCmd).Standalone()

	processes_upCmd.Flags().BoolP("detach", "d", false, "Start processes in the background")
	processes_upCmd.Flags().StringP("mode", "m", "", "The execution mode for process tasks (affects dependency resolution)")
	processes_upCmd.Flags().Bool("no-strict-ports", false, "Disable strict port mode, overriding strict_ports from devenv.yaml")
	processes_upCmd.Flags().Bool("strict-ports", false, "Error if a port is already in use instead of auto-allocating the next available port")

	processesCmd.AddCommand(processes_upCmd)

	processes_upCmd.MarkFlagsMutuallyExclusive("no-strict-ports", "strict-ports")

	carapace.Gen(processes_upCmd).FlagCompletion(carapace.ActionMap{
		"mode": actionExecutionModes(),
	})

	carapace.Gen(processes_upCmd).PositionalAnyCompletion(
		action.ActionProcesses(processes_upCmd).FilterArgs(),
	)
}
