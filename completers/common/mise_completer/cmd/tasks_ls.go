package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tasks_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List tasks",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasks_lsCmd).Standalone()

	tasks_lsCmd.Flags().BoolP("all", "a", false, "Show all tasks")
	tasks_lsCmd.Flags().BoolP("extended", "x", false, "Show extended information")
	tasks_lsCmd.Flags().Bool("hidden", false, "Show hidden tasks")
	tasks_lsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	tasksCmd.AddCommand(tasks_lsCmd)
}
