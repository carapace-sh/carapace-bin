package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var issue_closeCmd = &cobra.Command{
	Use:   "close [OPTIONS] WORK_ITEM_ID...",
	Short: "Close a specific WORK_ITEM_ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_closeCmd).Standalone()

	issue_closeCmd.Flags().Bool("help", false, "Show this message and exit.")
	issueCmd.AddCommand(issue_closeCmd)

	// TODO positional completion
}
