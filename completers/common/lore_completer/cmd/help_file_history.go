package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_file_historyCmd = &cobra.Command{
	Use:   "history",
	Short: "List revisions of a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_file_historyCmd).Standalone()

	help_fileCmd.AddCommand(help_file_historyCmd)
}
