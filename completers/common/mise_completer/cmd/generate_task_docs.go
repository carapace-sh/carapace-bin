package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generate_task_docsCmd = &cobra.Command{
	Use:   "task-docs",
	Short: "Generate task documentation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generate_task_docsCmd).Standalone()

	generateCmd.AddCommand(generate_task_docsCmd)
}
