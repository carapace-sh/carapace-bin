package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_threads_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List captured threads, newest first",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_threads_listCmd).Standalone()

	help_threadsCmd.AddCommand(help_threads_listCmd)
}
