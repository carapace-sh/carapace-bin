package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var todo_doneCmd = &cobra.Command{
	Use:   "done [<id>] [flags]",
	Short: "Mark a to-do item as done.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(todo_doneCmd).Standalone()

	todo_doneCmd.Flags().Bool("all", false, "Mark all pending to-do items as done.")
	todoCmd.AddCommand(todo_doneCmd)
}
