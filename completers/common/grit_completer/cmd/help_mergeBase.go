package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_mergeBaseCmd = &cobra.Command{
	Use:   "merge-base",
	Short: "Find merge bases or test whether one commit is an ancestor of another",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_mergeBaseCmd).Standalone()

	helpCmd.AddCommand(help_mergeBaseCmd)
}
