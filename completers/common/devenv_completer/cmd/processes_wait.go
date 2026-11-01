package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var processes_waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Wait for all processes to be ready",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processes_waitCmd).Standalone()

	processes_waitCmd.Flags().String("timeout", "", "Timeout in seconds")

	processesCmd.AddCommand(processes_waitCmd)
}
