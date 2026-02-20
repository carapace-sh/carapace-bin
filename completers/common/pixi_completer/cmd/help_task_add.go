package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_task_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a command to the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_task_addCmd).Standalone()

	help_taskCmd.AddCommand(help_task_addCmd)
}
