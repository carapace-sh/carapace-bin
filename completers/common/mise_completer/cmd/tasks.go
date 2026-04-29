package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:     "tasks",
	Aliases: []string{"t"},
	Short:   "List available tasks",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tasksCmd).Standalone()
	tasksCmd.Flags().Bool("json", false, "Output in JSON format")
	tasksCmd.Flags().Bool("extended", false, "Show extended task information")
	rootCmd.AddCommand(tasksCmd)
}
