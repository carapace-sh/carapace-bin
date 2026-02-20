package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var task_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a command to the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(task_help_addCmd).Standalone()

	task_helpCmd.AddCommand(task_help_addCmd)
}
