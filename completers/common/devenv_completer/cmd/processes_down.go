package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var processes_downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop processes running in the background",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(processes_downCmd).Standalone()

	processesCmd.AddCommand(processes_downCmd)
}
