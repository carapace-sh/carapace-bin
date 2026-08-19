package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var processes_attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach to running processes and stream their status and logs until Ctrl-C, leaving them running",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processes_attachCmd).Standalone()

	processesCmd.AddCommand(processes_attachCmd)
}
