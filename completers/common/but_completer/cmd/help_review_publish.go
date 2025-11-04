package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_review_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish review requests for active branches in your workspace. By default, publishes reviews for all active branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_review_publishCmd).Standalone()

	help_reviewCmd.AddCommand(help_review_publishCmd)
}
