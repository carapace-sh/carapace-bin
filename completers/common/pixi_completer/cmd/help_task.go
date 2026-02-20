package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Interact with tasks in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_taskCmd).Standalone()

	helpCmd.AddCommand(help_taskCmd)
}
