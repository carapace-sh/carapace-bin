package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var processes_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all managed processes and their status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processes_listCmd).Standalone()

	processesCmd.AddCommand(processes_listCmd)
}
