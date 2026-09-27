package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tasks_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show running and waiting tasks in concurrency groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_statusCmd).Standalone()

	tasks_statusCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	tasksCmd.AddCommand(tasks_statusCmd)
}
