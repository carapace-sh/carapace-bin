package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var open_issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Open all active issues view",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(open_issuesCmd).Standalone()

	open_issuesCmd.Flags().Bool("help", false, "Show this message and exit.")
	openCmd.AddCommand(open_issuesCmd)
}
