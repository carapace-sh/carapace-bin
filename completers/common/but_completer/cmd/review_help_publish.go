package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var review_help_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish review requests for active branches in your workspace. By default, publishes reviews for all active branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(review_help_publishCmd).Standalone()

	review_helpCmd.AddCommand(review_help_publishCmd)
}
