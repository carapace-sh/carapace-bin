package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var threads_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List captured threads, newest first",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(threads_help_listCmd).Standalone()

	threads_helpCmd.AddCommand(threads_help_listCmd)
}
