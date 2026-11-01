package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_help_historyCmd = &cobra.Command{
	Use:   "history",
	Short: "List revisions of a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_help_historyCmd).Standalone()

	file_helpCmd.AddCommand(file_help_historyCmd)
}
