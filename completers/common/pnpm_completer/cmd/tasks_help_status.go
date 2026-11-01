package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tasks_help_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show running and waiting tasks in concurrency groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_help_statusCmd).Standalone()

	tasks_helpCmd.AddCommand(tasks_help_statusCmd)
}
