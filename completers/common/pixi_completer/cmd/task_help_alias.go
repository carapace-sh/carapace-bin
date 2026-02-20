package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var task_help_aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Alias another specific command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(task_help_aliasCmd).Standalone()

	task_helpCmd.AddCommand(task_help_aliasCmd)
}
