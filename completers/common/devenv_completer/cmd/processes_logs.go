package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/devenv_completer/cmd/action"
	"github.com/spf13/cobra"
)

var processes_logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Get logs for a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processes_logsCmd).Standalone()

	processes_logsCmd.Flags().StringP("lines", "n", "", "Number of lines to show")
	processes_logsCmd.Flags().Bool("stderr", false, "Show only stderr")
	processes_logsCmd.Flags().Bool("stdout", false, "Show only stdout")

	processesCmd.AddCommand(processes_logsCmd)

	processes_logsCmd.MarkFlagsMutuallyExclusive("stderr", "stdout")

	carapace.Gen(processes_logsCmd).PositionalCompletion(
		action.ActionRunningProcesses(processes_logsCmd),
	)
}
