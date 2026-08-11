package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var file_historyCmd = &cobra.Command{
	Use:   "history",
	Short: "List revisions of a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_historyCmd).Standalone()

	file_historyCmd.Flags().String("branch", "", "Show branch revisions")
	file_historyCmd.Flags().String("depth", "", "Number of revisions to search initially")
	file_historyCmd.Flags().BoolP("help", "h", false, "Print help")
	file_historyCmd.Flags().Bool("oneline", false, "Output each revision on one line only")
	file_historyCmd.Flags().String("revision", "", "Revision to start from")
	fileCmd.AddCommand(file_historyCmd)
}
