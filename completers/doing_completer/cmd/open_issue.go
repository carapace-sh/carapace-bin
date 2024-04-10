package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var open_issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Open a specific WORK_ITEM_ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(open_issueCmd).Standalone()

	open_issueCmd.Flags().Bool("help", false, "Show this message and exit.")
	openCmd.AddCommand(open_issueCmd)
}
