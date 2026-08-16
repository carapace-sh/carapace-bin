package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tasks_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available tasks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_listCmd).Standalone()

	tasks_listCmd.Flags().Bool("json", false, "Print tasks as JSON for machine consumption")

	tasksCmd.AddCommand(tasks_listCmd)
}
