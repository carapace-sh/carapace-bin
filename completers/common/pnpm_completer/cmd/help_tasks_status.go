package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_tasks_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show running and waiting tasks in concurrency groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_tasks_statusCmd).Standalone()

	help_tasksCmd.AddCommand(help_tasks_statusCmd)
}
