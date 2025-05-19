package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pr_closeCmd = &cobra.Command{
	Use:   "close [OPTIONS] PR_ID...",
	Short: "Close a specific PR_ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_closeCmd).Standalone()

	pr_closeCmd.Flags().Bool("help", false, "Show this message and exit.")
	prCmd.AddCommand(pr_closeCmd)
}
