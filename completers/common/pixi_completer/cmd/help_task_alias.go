package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_task_aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Alias another specific command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_task_aliasCmd).Standalone()

	help_taskCmd.AddCommand(help_task_aliasCmd)
}
